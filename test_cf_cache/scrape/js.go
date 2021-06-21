package scrape

import (
	"github.com/PuerkitoBio/goquery"
)

func GetJs(document *goquery.Document) []string {
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
