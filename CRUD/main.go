package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct{
	UserID int `json:"user_id"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func performGetreq(){
	res,err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err !=nil{
        fmt.Println("Error making request:", err)
        return
    }
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK{
		fmt.Printf("Error getting data, status code: %d\n", res.StatusCode)
        return
	}

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil{
		fmt.Println("Error decoding JSON:", err)
        return
	}
	// fmt.Println("Todo: ",todo)
	fmt.Printf("Todo: %+v\n", todo)
}

func performPostreq(){
	todo := Todo{
		UserID: 7,
        Title: "Kunal Koshta",
        Completed: true,
	}

	jsonData, err := json.Marshal(todo)
	if err !=nil{
        fmt.Println("Error marshalling JSON:", err)
        return
    }

	jsonStr := string(jsonData)

	jsonReader := strings.NewReader(jsonStr)
	myURl := "https://jsonplaceholder.typicode.com/todos"
	res, err := http.Post(myURl,"application/json", jsonReader)
	if err != nil{
        fmt.Println("Error making POST request:", err)
        return
    }
	defer res.Body.Close()

	data,err := io.ReadAll(res.Body)
	if err !=nil{
        fmt.Println("Error reading response body:", err)
        return
    }
	fmt.Println("Response body:", string(data))
	fmt.Println("Status code:", res.Status)
}

func performUpdate(){
	todo := Todo{
        UserID: 7,
        Id: 101,
        Title: "Kunal Koshta Updated",
        Completed: false,
    }

    jsonData, err := json.Marshal(todo)
    if err !=nil{
        fmt.Println("Error marshalling JSON:", err)
        return
    }

    jsonStr := string(jsonData)
    jsonReader := strings.NewReader(jsonStr)

    myURL := "https://jsonplaceholder.typicode.com/todos/1"
    req, err := http.NewRequest(http.MethodPut, myURL, jsonReader)
    if err != nil{
        fmt.Println("Error making PUT request:", err)
        return
    }
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil{
        fmt.Println("Error making PUT request:", err)
        return
    }

    data,err := io.ReadAll(res.Body)
	if err !=nil{
        fmt.Println("Error reading response body:", err)
        return
    }
	defer res.Body.Close()
	fmt.Println("Dat: ",string(data))
	fmt.Println("Status code: ", res.Status)
}

func performDelete(){
	myURl := "https://jsonplaceholder.typicode.com/todos/1"
	req, err := http.NewRequest(http.MethodDelete, myURl, nil)
	if err != nil{
        fmt.Println("Error making DELETE request:", err)
        return
    }
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil{
        fmt.Println("Error making DELETE request:", err)
        return
    }
	defer res.Body.Close()
	fmt.Println("Status code: ", res.Status)
}

func main()  {
	fmt.Println("CRUD")
	// performGetreq()
	// performPostreq()
	// performUpdate()
	performDelete()
}