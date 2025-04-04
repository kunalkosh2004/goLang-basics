package main

import "fmt"

func main()  {
	fmt.Println("Maps")

	languages := make(map[string]string)

	languages["PY"] = "python"
	languages["GO"] = "golang"
	languages["JS"] = "javascript"
	languages["RB"] = "ruby"

	fmt.Println("Languages are: ",languages)
	fmt.Println("GO stands for: ",languages["GO"])

	// Deleting a key-value pair
	delete(languages, "JS")
	fmt.Println("After deleting a key-value pair: ",languages)
}