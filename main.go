package main

import "fmt"

// main function prints "Hello, World!" and finds the sum of two numbers entered by the user.
// It prompts the user to enter two numbers and then calculates their sum using the Scanln function from the fmt package.
func main() {
	// code block
}
func main() {
	fmt.Println("Hello, World!")
	//find sum of two numbers
	var a, b int
	fmt.Println("Enter two numbers")

	fmt.Scanln(&a)
	fmt.Scanln(&b)
	
	fmt.Println("Sum of two numbers is ", a+b)
	
}
