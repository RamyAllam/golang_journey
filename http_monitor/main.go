package main

import (
	"fmt"
	"log"
	"time"

	"go.uber.org/config"

	"github.com/RamyAllam/golang_journey/http_monitor/monitor"
	"github.com/slack-go/slack"
)

type urlCfg []string

type globalCfg struct {
	ValidHttpCode []int         `yaml:"valid_http_code"`
	HttpTimeout   time.Duration `yaml:"http_connection_timeout"`
	Slack         bool          `yaml:"slack"`
	SlackHook     string        `yaml:"slack_webhook"`
}

var urls urlCfg
var gcfg globalCfg

func init() {
	// Construct the YAML
	provider, err := config.NewYAML(config.File("config.yaml"))

	if err != nil {
		panic(err)
	}

	// Populate the yaml file to the urls slice
	if err := provider.Get("url").Populate(&urls); err != nil {
		panic(err)
	}

	// Populate the yaml file to the global config
	if err := provider.Get("global").Populate(&gcfg); err != nil {
		panic(err)
	}

	if gcfg.Slack {
		if len(gcfg.SlackHook) < 10 {
			log.Fatal("Slack webhook URL is empty")
		}
	}
}

func main() {

	// Loop through the urls slice and initiate the monitor
	for _, url := range urls {
		statusCode, statusErr := monitor.GetStatusCode(url, gcfg.HttpTimeout)

		if statusErr != nil {
			statusMsg := fmt.Sprintf("%s:\t%s", url, statusErr)
			log.Println(statusMsg)

			/*
				Check if Slack alert is enabled
				 Alert only when the HTTP is reporting an error
			*/
			if gcfg.Slack {
				slackErr := slack.PostWebhook(gcfg.SlackHook, &slack.WebhookMessage{Text: statusMsg})

				if slackErr != nil {
					log.Println(slackErr)
				}
			}

			// If there are no errors, check for the HTTP status code
		} else {

			// Check if the HTTP status code is in the valid HTTP status code list
			_, validCode := findSlice(gcfg.ValidHttpCode, statusCode)

			// If the HTTP code is is not found healthy, print and alert
			if !validCode {
				statusMsg := fmt.Sprintf("%s:\tDown with status code: %d", url, statusCode)
				log.Println(statusMsg)

				/*

					Check if Slack alert is enabled
					Alert only when the site is down
				*/
				if gcfg.Slack {
					slackErr := slack.PostWebhook(gcfg.SlackHook, &slack.WebhookMessage{Text: statusMsg})

					if slackErr != nil {
						log.Println(slackErr)
					}
				}

				// Log when the site is Up
			} else {
				statusMsg := fmt.Sprintf("%s:\tUp with status code: %d", url, statusCode)
				log.Println(statusMsg)
			}
		}

	}

}
