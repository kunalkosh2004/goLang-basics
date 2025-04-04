package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("File Handling")
	/*
	file, err := os.Create("example.txt")

	if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
	defer file.Close()

	content := "Hello by Kunal!"

	_, err = io.WriteString(file, content+"\n")

	if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }
	fmt.Println("File created successfully!")
	*/

	/*
	// Reading from a file in Go
    // This will read the file line by line and print them on the console
    // It will stop reading when it encounters an error or reaches the end of the file
    //defer is used to close the file after the function execution is done
    //ioutil.ReadAll is used to read the entire file content into a byte slice
    //ioutil.ReadAll will return an error if the file does not exist or if there is an error reading the file
    //fmt.Println(ioutil.ReadAll(file))

    // Using ReadFile function directly
    /*
    data, err := ioutil.ReadFile("example.txt")

    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    fmt.Println(string(data))
    */
	
	/*
    // Using Read function to read data line by line
	file, err := os.Open("example.txt")

	if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
	defer file.Close()

	buffer := make([]byte,1024)

	for {
		n,err := file.Read(buffer)
		if err == io.EOF {
            break
        }
		if err != nil {
            fmt.Println("Error reading from file:", err)
            return
        }

		fmt.Println(string(buffer[:n]))
	}
		*/

	// Not goo for large file as it loads the whole data at once in the memory
	content,err := os.ReadFile("example.txt")

	if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

	fmt.Println(string(content))
}