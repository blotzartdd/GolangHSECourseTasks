package main

import (
	"fmt"
)

func check(num int) bool {
	return num%2 == 0
}

func main() {
	var num int

	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println(err)
		return
	}

	if check(num) {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
