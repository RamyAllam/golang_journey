package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/RamyAllam/golang_journey/coinsapp_yamlconfig/api"
	"github.com/RamyAllam/golang_journey/coinsapp_yamlconfig/slack"
	"go.uber.org/config"
)

type cfg struct {
	ApiKey    string `yaml:"api_key"`
	Slack     bool   `yaml:"slack"`
	SlackHook string `yaml:"slack_webhook"`
}

func main() {

	fiatCurrency := flag.String(
		"fiat", "USD", "The name of the fiat currency",
	)

	cryptoCurrency := flag.String(
		"crypto", "BTC", "The crypto currency name",
	)

	flag.Parse()

	// Construct the YAML
	provider, err := config.NewYAML(config.File("config.yaml"))

	if err != nil {
		panic(err)
	}

	// Populate the yaml file to the struct
	var c cfg
	if err := provider.Get("global").Populate(&c); err != nil {
		panic(err)
	}

	// Pass the variables to FetchCrypto and perform the API call
	crypto, err := api.FetchCrypto(*fiatCurrency, *cryptoCurrency, c.ApiKey)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(crypto)

	// Send to slack when the slack is set in the yaml file
	if c.Slack {
		fmt.Println("Sending a message to slack")
		err = slack.SendMessage(c.SlackHook, crypto)
		if err != nil {
			log.Fatal(err)
		}
	}
}
