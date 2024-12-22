package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello")
	fmt.Println("Please rate ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for ,", input)

	fmt.Printf("type of input is %T \n", input)

	numrat, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)

	fmt.Println("The converted and incremented rating is ", numrat+1)

}
