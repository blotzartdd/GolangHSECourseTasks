package main

import (
	"fmt"
)

func isVowel(char string) bool {
	switch char {
	case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
		return true
	default:
		return false
	}
}

func main() {
	var s string

	_, err := fmt.Scan(&s)
	if err != nil {
		fmt.Println(err)
		return
	}

	if isVowel(s) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
