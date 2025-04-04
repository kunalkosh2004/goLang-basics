package main

import "fmt"

func main() {
	var username string = "Kunal Koshta"
	fmt.Println(username)
	fmt.Printf("Variables is of type: %T \n",username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var small_val uint8 = 255
	fmt.Println(small_val)
	fmt.Printf("variable is of type: %T \n", small_val)

	var small_float float64 = 255.2352435345245
	fmt.Println(small_float)
	fmt.Printf("variable is of type: %T \n", small_float)


	// default values and aliases
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("variable is of type: %T \n", anotherVar)

	var website = "www.example.com"
	fmt.Println(website)

	numberOfUsers := 123
	fmt.Println(numberOfUsers)
	fmt.Printf("Type: %T \n", numberOfUsers)
}	

