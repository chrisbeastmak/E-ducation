package main

import "fmt"

func main() {
	// Welcome to Go Tutorial
	fmt.Printf("Welcome to the E-ducation for Go Programming!\n")
 
	// Variables
	// Declaring variable
	var age int
	age = 29
	fmt.Printf("My age is %d\n", age)
 
	// String Variable
	var name string   
	name = "John Smith"  
	fmt.Printf("My name is %s\n", name)
 
	// Float Variable 
	var gpa float32
	gpa = 3.50
	fmt.Printf("My GPA is %f\n", gpa)
 
	// Boolean Variable 
	var isValid bool
	isValid = true
	fmt.Printf("Is this true? %t\n", isValid)
 
	// Arrays
	// Array of Integer
	var numbers = [3]int{1, 2, 3}
	fmt.Printf("My array of integers is %v\n", numbers)
 
	// Array of Strings
	var names = [3]string{"John", "Adam", "David"}
	fmt.Printf("My array of strings is %v\n", names)
 
	// Array of Boolean
	var results = [3]bool{true, false, true}
	fmt.Printf("My array of boolean is %v\n", results)
 
	// Slices
	// Slice of Integer 
	var numberSlice = []int{1, 2, 3, 4, 5}
	fmt.Printf("My slice of integers is %v\n", numberSlice)
 
	// Slice of Strings 
	var namesSlice = []string{"John", "Adam", "David", "Michael"}
	fmt.Printf("My slice of strings is %v\n", namesSlice)
 
	// Slice of Boolean
	var resultSlice = []bool{true, false, true, false}
	fmt.Printf("My slice of boolean is %v\n", resultSlice)
 
	// Maps
	// Map of Integer 
	var scores = map[string]int{"John": 10, "Adam": 15, "David": 20}
	fmt.Printf("My map of integers is %v\n", scores)
 
	// Map of Strings
	var grades = map[string]string{"John": "A", "Adam": "B", "David": "C"}
	fmt.Printf("My map of strings is %v\n", grades)
 
	// Map of Boolean
	var records = map[string]bool{"John": true, "Adam": false, "David": true}
	fmt.Printf("My map of boolean is %v\n", records)
 
	// Conditionals
	// If Statement
	if age >= 18 {
		fmt.Printf("You are an adult\n")
	} else {
		fmt.Printf("You are not an adult\n")
	}
 
	// Switch Statement
	switch results[0] {
	case true:
		fmt.Printf("Result is true\n")
	case false:
		fmt.Printf("Result is false\n")
	}
 
	// Loops
	// For Loop 
	for i := 0; i < len(numberSlice); i++ {
		fmt.Printf("%d\n", numberSlice[i])
	}
 
	// For Range Loop 
	for _, name := range namesSlice {
		fmt.Printf("%s\n", name)
	}
 
	// While Loop
	i := 0
	while i < len(resultSlice) {
		fmt.Printf("%t\n", resultSlice[i])
		i++
	}
 
	// Functions
	// Declare function 
	func addNumber(num1 int, num2 int) int {
		return num1 + num2
	}
 
	// Calling function
	sum := addNumber(1, 2)
	fmt.Printf("Sum is %d\n", sum)
 
	// Structures 
	// Declaring structure
	type Student struct {
		name    string
		age     int
		gpa     float32
		isValid bool
	}
 
	// Initializing structure
	student1 := Student{"John", 29, 3.50, true}
	fmt.Printf("Student1 : %v\n", student1)
 
	// Pointer
	// Create pointer
	counter := 0
	counterPointer := &counter
 
	// Read value from pointer
	fmt.Printf("Counter value is %d\n", *counterPointer)
 
	// Edit value from pointer
	*counterPointer = 2
	fmt.Printf("Counter value is %d\n", *counterPointer)
 
	// Conclusion
	fmt.Printf("You now understand the basics of Go Programming!\n")
}