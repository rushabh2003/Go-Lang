package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting Maps")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["Py"] = "Python"

	fmt.Println("List of all languages ", languages)
	fmt.Println("JS is short for", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages ", languages)

	//loops preview

	for key, value := range languages {
		fmt.Printf("For key %v , value is %v\n", key, value)
	}
}
