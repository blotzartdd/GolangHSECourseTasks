package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	_, err := fmt.Scan(&a, &b, &c)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(max(a, max(b, c)))
}
