package main

import "fmt"

func main()  {
	fmt.Println("Slices")
	
	// var courses = []string{}
	// courses = append(courses, "Go")
	// courses = append(courses, "Python")
	// courses = append(courses, "JavaScript")
	// courses = append(courses, "Ruby")
	// courses = append(courses, "C++")

    // fmt.Println(courses)
	// // fmt.Println(len(courses))
	// // removing a value from slice using index
	// var index int = 2
	// courses = append(courses[:index], courses[index+1:]...)
	// fmt.Println(courses)

	numbers := make([]int,0)

	numbers = append(numbers,1)
	numbers = append(numbers,2)
	numbers = append(numbers,3)
	numbers = append(numbers,4)
	numbers = append(numbers,5)
	fmt.Println("Slice: ",numbers)
	fmt.Println("Length: ",len(numbers))
	fmt.Println("Capacity: ",cap(numbers))
}