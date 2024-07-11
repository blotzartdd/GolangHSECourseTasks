package main

import (
	"fmt"
)

func getFact(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}

	return res
}

func main() {
	var n int

	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(getFact(n))
}
