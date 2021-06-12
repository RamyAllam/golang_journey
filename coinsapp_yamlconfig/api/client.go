package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func FetchCrypto(fiat string, crypto string, apikey string) (string, error) {
	URL := "https://api.nomics.com/v1/currencies/ticker?key=" + apikey + "&interval=1d&ids=" + crypto + "&convert=" + fiat

	resp, err := http.Get(URL)

	if err != nil {
		log.Fatal("API HTTP error happened\n", err)
	}

	defer resp.Body.Close()

	var cResponse cryptoResponse

	if err := json.NewDecoder(resp.Body).Decode(&cResponse); err != nil {
		log.Fatal("Error while encoding the response")
	}

	return cResponse.TextOutput(), nil
}
