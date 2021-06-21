package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func (s *site) document() (*goquery.Document, error) {
	// Make HTTP Request
	response, err := http.Get(s.url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	// Create a goquery document from the HTTP Response
	document, err := goquery.NewDocumentFromReader(response.Body)

	return document, err
}

func (s *site) images(document *goquery.Document) []string {
	var urls []string

	// Find tags
	document.Find("img").Each(func(i int, element *goquery.Selection) {
		// Get the src value
		value, exists := element.Attr("src")
		if exists {
			urls = append(urls, value)
		}
	})

	return urls
}

func (s *site) css(document *goquery.Document) []string {
	var urls []string

	// Find URLs
	document.Find("link").Each(func(i int, element *goquery.Selection) {
		// Get the rel attribute
		value, exists := element.Attr("rel")

		if exists {
			// Check if the rel is stylesheet
			if value == "stylesheet" {

				// Get the href value
				value, exists = element.Attr("href")

				if exists {
					urls = append(urls, value)
				}
			}
		}
	})

	return urls
}

func (s *site) js(document *goquery.Document) []string {
	var urls []string

	// Find tags
	document.Find("script").Each(func(i int, element *goquery.Selection) {

		// Get the src value
		value, exists := element.Attr("src")

		if exists {
			urls = append(urls, value)
		}
	})

	return urls
}
