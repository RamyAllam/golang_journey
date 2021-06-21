package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/RamyAllam/golang_journey/test_cf_cache/cloudflare"
	"github.com/RamyAllam/golang_journey/test_cf_cache/scrape"
)

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

	// Create a goquery document from the HTTP Response
	document, err := scrape.GetDocument(*siteUrl)

	if err != nil {
		log.Fatal("Error loading HTTP Response Body", err)
	}

	if (*assetType) == "images" || (*assetType) == "all" {
		fmt.Println()
		imagesList := generateImagesList(document, siteUrl)

		for _, v := range imagesList {
			fmt.Println("Testing URL: ", v)
			cloudflare.Report(v)
			fmt.Println("----------------------")
		}

	}

	if (*assetType) == "css" || (*assetType) == "all" {
		fmt.Println()

		cssList := generateCSSList(document, siteUrl)
		for _, v := range cssList {
			fmt.Println("Testing URL: ", v)
			cloudflare.Report(v)
			fmt.Println("----------------------")
		}
	}

	if (*assetType) == "js" || (*assetType) == "all" {
		fmt.Println()

		jsList := generateJSList(document, siteUrl)

		for _, v := range jsList {
			fmt.Println("Testing URL: ", v)
			cloudflare.Report(v)
			fmt.Println("----------------------")
		}
	}

}
