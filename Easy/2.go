package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func main() {
	var a int
	var b int

	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(add(a, b))
}
