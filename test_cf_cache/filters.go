package main

import (
	"fmt"
	"net/url"
	"strings"
)

func filterUrls(srcUrl string, destUrls []string) []string {
	/*
		This function checks the list of the scrapped URLs
		and making sure they are all of the source site and not external ones
		Params:
			- srcUrl: The URL that the user specifies (string)
			- srcUrl: The list of the scrapped URLs ([]string)
		Terms:
			- Source URL: The URL that the user specifies
			- Dest URL: The URLs we gather from the scrapping process
	*/
	var results []string

	/*
		Parse the source URL
	*/
	srcUrlParsed, err := url.Parse(srcUrl)
	if err != nil {
		panic(err)
	}

	// Get the host attribute of the source URL
	srcHost := srcUrlParsed.Host
	srcScheme := srcUrlParsed.Scheme

	/*
		Parse the destination URL
	*/
	// Loop through the slice of the dest URLs
	for _, v := range destUrls {
		vDestParsed, err := url.Parse(v)
		if err != nil {
			panic(err)
		}

		// Get the host attribute of the destination URL
		// Check if the values are all for the same site, not external site
		destHost := vDestParsed.Host

		if srcHost == destHost {
			results = append(results, v)
		}

		//  Skip values for external URLs and a mixed content
		//  Ex. //cdn-images.mailchimp.com/embedcode/classic-10_7.css
		if strings.HasPrefix(v, "//") {
			continue
		}

		// Handle Relative URLs
		// Ex. css/modern.css
		if !strings.HasPrefix(v, "http") {
			fullRelativeUrl := fmt.Sprintf("%s://%s/%s", srcScheme, srcHost, v)
			results = append(results, fullRelativeUrl)
		}

	}

	return results
}

func (s *site) filterImagesList(assets []string, siteUrl string) []string {
	assetsList := filterUrls(siteUrl, assets)

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

func (s *site) filterCSSList(assets []string, siteUrl string) []string {
	assetsList := filterUrls(siteUrl, assets)

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

func (s *site) filterJSList(assets []string, siteUrl string) []string {
	assetsList := filterUrls(siteUrl, assets)

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
