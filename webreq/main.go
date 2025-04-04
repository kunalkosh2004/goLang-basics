package main

import (
	"fmt"
	"io"
	"net/http"
)

func main()  {
	fmt.Println("Web request")
	
	res,err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error making request:", err)
        return
	}
	defer res.Body.Close()
	fmt.Printf("Type of the res %T:\n", res)

	data, err := io.ReadAll(res.Body)
	if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    fmt.Println("Response body:", string(data))
}