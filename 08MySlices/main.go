package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Let's start slices")

	var fruits = []string{"Apple ", " Banana", "Peach"}

	fmt.Printf("Type of fruits list is %T\n", fruits)

	fruits = append(fruits, "Kiwi")

	fmt.Println(fruits)

	fruits = append(fruits[:3])

	fmt.Println(fruits)

	highscore := make([]int, 4)

	highscore[0] = 234
	highscore[1] = 333
	highscore[2] = 299
	highscore[3] = 655
	// highscore[4] = 455

	highscore = append(highscore, 466, 599)

	fmt.Println(highscore)

	sort.Ints(highscore)

	fmt.Println(highscore)
	fmt.Println(sort.IntsAreSorted(highscore))

	//Removing an element using append

	var colors = []string{"Red", "Yellow", "Blue", "Green", "Pink"}

	fmt.Println(colors)

	var index int = 2

	colors = append(colors[:index], colors[index+1:]...)

	fmt.Println("Slice after deleting 2nd index value ", colors)

}
