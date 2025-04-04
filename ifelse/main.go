package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("If Else Condition")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your age: ")
	age, _ := reader.ReadString('\n')
	numAge, err := strconv.ParseInt(strings.TrimSpace(age),16,64)
	fmt.Println("Age: ", numAge)
	if err != nil {
		fmt.Println("Error reading input:", err)
        return
	} else{
		if numAge >= 18 {
            fmt.Println("You are an adult")
        } else{
            fmt.Println("You are a child")
        }
	}

	if num := 4; num<10 {
		fmt.Println("Number is less than 10")
	} else{
		fmt.Println("Number is greater than or equal to 10")
	}
}