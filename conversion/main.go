package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("Welcome to my Pizz app")
	fmt.Println("Please rate pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your rating: ")
	rating, _ := reader.ReadString('\n')

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating),64)

	if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }else{
		fmt.Println("Thanks for rating: ", numRating+1)
	}
}