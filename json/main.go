package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int `json:"age"`
	IsAdult bool `json:"is-adult"`
}

func main()  {
	fmt.Println("Learning JSON")

	person := Person{Name:"Kunal",Age: 21,IsAdult: true}

	fmt.Println("Person Data: ",person)
	
	// Encoding Marshal	JSON
	jsonData,err := json.Marshal(person)
	if err !=nil{
        fmt.Println("Error marshalling JSON:", err)
        return
    }

	fmt.Println("JSON Data: ", string(jsonData))

	// Decoding Unmarshal JSON
	var personData Person

	err = json.Unmarshal(jsonData, &personData)
	if err !=nil{
        fmt.Println("Error unmarshalling JSON:", err)
        return
    }
	fmt.Println("Decoded Person Data: ", personData)
}