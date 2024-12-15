package main

import "fmt"

const LoginToken = "arnold" //public

func main() {
	var username string = "Rushabh"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallVAl uint8 = 255
	fmt.Println(smallVAl)
	fmt.Printf("Variable is of type : %T \n", smallVAl)

	var smallFloat float32 = 255.1126232
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	//default values

	var anotherVariable string
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n ", anotherVariable)

	//implicit type

	var website = "mywebsite.in"
	fmt.Println(website)

	//no var style

	numberOfUsers := 1729.23
	//should not be used outside methods
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n ", LoginToken)

}
