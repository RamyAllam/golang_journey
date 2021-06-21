package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/RamyAllam/golang_journey/test_cf_cache/scrape"
)

func generateImagesList(document *goquery.Document, siteUrl *string) []string {
	assets := scrape.GetImages(document)
	assetsList := filterUrls(*siteUrl, assets)

	if len(assetsList) > 0 {
		fmt.Println("Images List:")

		for i, v := range assetsList {
			i += 1
			fmt.Printf("%d) %s\n", i, v)
		}
	} else {
		fmt.Println("Images List: Empty")
	}

	return assetsList
}

func generateCSSList(document *goquery.Document, siteUrl *string) []string {
	assets := scrape.GetCss(document)
	assetsList := filterUrls(*siteUrl, assets)

	if len(assetsList) > 0 {
		fmt.Println("CSS List:")

		for i, v := range assetsList {
			i += 1
			fmt.Printf("%d) %s\n", i, v)
		}
	} else {
		fmt.Println("CSS List: Empty")
	}

	return assetsList
}

func generateJSList(document *goquery.Document, siteUrl *string) []string {
	assets := scrape.GetJs(document)
	assetsList := filterUrls(*siteUrl, assets)

	if len(assetsList) > 0 {
		fmt.Println("JS List:")

		for i, v := range assetsList {
			i += 1
			fmt.Printf("%d) %s\n", i, v)
		}
	} else {
		fmt.Println("JS List: Empty")
	}

	return assetsList
}
