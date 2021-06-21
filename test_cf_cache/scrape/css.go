package scrape

import (
	"github.com/PuerkitoBio/goquery"
)

func GetCss(document *goquery.Document) []string {
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
