package scrape

import (
	"github.com/PuerkitoBio/goquery"
)

func GetImages(document *goquery.Document) []string {
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
