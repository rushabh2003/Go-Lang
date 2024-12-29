package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNum := 25

	var ptr = &myNum
	fmt.Println("Value of actual poinnter is ", ptr)
	fmt.Println("Value of actual poinnter is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value is : ", myNum)
}
