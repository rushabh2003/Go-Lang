package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello")

	presenttime := time.Now()
	fmt.Println(presenttime)

	fmt.Println(presenttime.Format("01-02-2006 15:04:05 Monday")) //standard format

	createDate := time.Date(2020, time.August, 13, 23, 23, 0, 0, time.UTC)

	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))

}
