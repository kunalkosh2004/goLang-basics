package main

import "fmt"

func main()  {
	fmt.Println("Welcome to Methods")
	kunal := Person{"Kunal", 21, "CSE", "Male"}
	fmt.Println(kunal)
	fmt.Printf("Kunal's details are: %+v\n", kunal)
}

type Person struct {
    Name string
    Age  int
	Branch    string
	Gender    string
}