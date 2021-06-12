package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/RamyAllam/golang_journey/coinsapp/api"
	"github.com/RamyAllam/golang_journey/coinsapp/slack"
)

func main() {

	fiatCurrency := flag.String(
		"fiat", "USD", "The name of the fiat currency",
	)

	cryptoCurrency := flag.String(
		"crypto", "BTC", "The crypto currency name",
	)

	apiKey := flag.String(
		"apikey", "", "Nomics API key",
	)

	slackWebhook := flag.String(
		"slack", "", "Slack webhook URL",
	)

	flag.Parse()

	// Make sure Nomics API Key is passed
	if len(*apiKey) < 10 {
		log.Fatal("Nomics API Key is empty")
	}

	crypto, err := api.FetchCrypto(*fiatCurrency, *cryptoCurrency, *apiKey)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(crypto)

	// Slack is optional
	if len(*slackWebhook) > 0 {
		err = slack.SendMessage(*slackWebhook, crypto)
		if err != nil {
			log.Fatal(err)
		}
	}

}
