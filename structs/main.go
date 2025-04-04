package main

import "fmt"

func main()  {
	fmt.Println("Structs")
	
	kunal := Person{"Kunal", 21, "CSE", "Male"}

	fmt.Printf("Details of kunal are: %+v\n", kunal)
	fmt.Printf("Name is: %v and Branch is: %v\n", kunal.Name, kunal.Branch)
}

type Person struct {
    Name string
    Age  int
	Branch	string
	Gender    string
}
