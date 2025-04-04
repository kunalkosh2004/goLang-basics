package main

import "fmt"

func main()  {
	fmt.Println("Welcome to Functions")
	greeter()

	result := adder(2,3)
	fmt.Printf("Result: %d\n", result)
}

func adder(a int, b int) int {
    return a + b
}

func greeter(){
	fmt.Println("Namaste from Kunal...")
}