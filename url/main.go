package main

import (
	"fmt"
	"net/url"
)

func main()  {
	fmt.Println("URLs")
	myURL := "http://example.com:8080/path/to/resource?key1=value1&key2=value2"

	fmt.Println("My URL: ", myURL)
	fmt.Printf("Type of URL: %T \n", myURL)

	parsedURL,err := url.Parse(myURL)

	if err!=nil{
		fmt.Println("Error parsing URL:", err)
        return
	}

	fmt.Printf("Type of ParsedURL: %T \n", parsedURL)

	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Raw Query:", parsedURL.RawQuery)

	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "username=kunalkosh"

	newURL := parsedURL.String()

	fmt.Println("New URL: ", newURL)

}