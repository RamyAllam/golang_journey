package cloudflare

import "fmt"

func Report(siteUrl string) {
	headers, code := GetResponseData(siteUrl, 10)

	fmt.Println("Status Code:", code)
	fmt.Println("Cf-Cache-Status:", headers.Get("Cf-Cache-Status"))
	fmt.Println("Cache-Control:", headers.Get("Cache-Control"))
	fmt.Println("ki-Cache-Tag:", headers.Get("ki-Cache-Tag"))
	fmt.Println("ki-edge:", headers.Get("ki-edge"))
	fmt.Println("X-Edge-Location-Klb:", headers.Get("X-Edge-Location-Klb"))
	fmt.Println("CF-Ray:", headers.Get("CF-Ray"))
}
