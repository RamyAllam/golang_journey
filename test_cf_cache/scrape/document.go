package scrape

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetDocument(url string) (*goquery.Document, error) {

	// Make HTTP Request
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	// Create a goquery document from the HTTP Response
	document, err := goquery.NewDocumentFromReader(response.Body)

	return document, err
}
