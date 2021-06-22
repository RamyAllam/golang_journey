package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/RamyAllam/golang_journey/test_cf_cache/cloudflare"
)

type site struct {
	url string
}

func main() {

	banner :=
		`
		_____ ______   _____            _            _____         _            
		/  __ \|  ___| /  __ \          | |          |_   _|       | |           
		| /  \/| |_    | /  \/ __ _  ___| |__   ___    | | ___  ___| |_ ___ _ __ 
		| |    |  _|   | |    / _- |/ __| '_ \ / _ \   | |/ _ \/ __| __/ _ \ '__|
		| \__/\| |     | \__/\ (_| | (__| | | |  __/   | |  __/\__ \ ||  __/ |   
		 \____/\_|      \____/\__,_|\___|_| |_|\___|   \_/\___||___/\__\___|_|   
		`

	fmt.Println(banner)

	assetType := flag.String(
		"asset", "all", "The type of the assets",
	)

	siteUrl := flag.String(
		"url", "", "The site URL",
	)

	flag.Parse()

	if len(*siteUrl) < 5 {
		log.Fatal("Please enter a valid URL")
	}

	inputSite := site{
		url: *siteUrl,
	}

	siteDocument, err := inputSite.document()

	if err != nil {
		log.Fatal(err)
	}

	if (*assetType) == "images" || (*assetType) == "all" {
		siteDocumentAssets := inputSite.images(siteDocument)
		filteredAssetsList := filterUrls(inputSite.url, siteDocumentAssets)

		if len(filteredAssetsList) > 0 {
			fmt.Println("Images List:")

			for i, v := range filteredAssetsList {
				i += 1
				fmt.Printf("%d) %s\n", i, v)
			}
		} else {
			fmt.Println("Images List: Empty")
		}

		for _, v := range filteredAssetsList {
			fmt.Println("Testing URL: ", v)
			cloudflare.Report(v)
			fmt.Println("----------------------")
		}

	}

	if (*assetType) == "css" || (*assetType) == "all" {
		siteDocumentAssets := inputSite.css(siteDocument)
		filteredAssetsList := filterUrls(inputSite.url, siteDocumentAssets)

		if len(filteredAssetsList) > 0 {
			fmt.Println("CSS List:")

			for i, v := range filteredAssetsList {
				i += 1
				fmt.Printf("%d) %s\n", i, v)
			}
		} else {
			fmt.Println("CSS List: Empty")
		}

		for _, v := range filteredAssetsList {
			fmt.Println("Testing URL: ", v)
			cloudflare.Report(v)
			fmt.Println("----------------------")
		}

	}

	if (*assetType) == "js" || (*assetType) == "all" {
		siteDocumentAssets := inputSite.js(siteDocument)
		filteredAssetsList := filterUrls(inputSite.url, siteDocumentAssets)

		if len(filteredAssetsList) > 0 {
			fmt.Println("JS List:")

			for i, v := range filteredAssetsList {
				i += 1
				fmt.Printf("%d) %s\n", i, v)
			}
		} else {
			fmt.Println("JS List: Empty")
		}

		for _, v := range filteredAssetsList {
			fmt.Println("Testing URL: ", v)
			cloudflare.Report(v)
			fmt.Println("----------------------")
		}

	}

}
