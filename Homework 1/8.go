package main

import (
	"fmt"
)

func reverse(s string) (res string) {
	for _, char := range s {
		res = string(char) + res
	}

	return res
}

func main() {
	var s string

	_, err := fmt.Scan(&s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(reverse(s))
}
