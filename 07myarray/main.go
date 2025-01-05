package main

import "fmt"

func main() {
	fmt.Println("Let's do arrays")

	//Arrays are not extensievly used in Go Lang.

	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[3] = "Kiwi"

	fmt.Println("fruit list is ", fruits)
	fmt.Println("length of array ", len(fruits))

	var vegies = [3]string{"Bhindi", "Karela", "Palak"}

	fmt.Println("vegies are ", vegies)
}
