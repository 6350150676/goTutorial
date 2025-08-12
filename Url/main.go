package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL")

	myurl := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"

	fmt.Printf("Type of URL: %T\n", myurl)

	parsedUrl, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("Raw Query:", parsedUrl.RawQuery)

	parsedUrl.Path = "/newPath"
	parsedUrl.RawQuery = "username=jhatkade"

	newUrl := parsedUrl.String()

	fmt.Println("new URL: ", newUrl)
}
