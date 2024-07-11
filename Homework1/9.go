package main

import (
	"fmt"
)

func sum(arr []int) (sum int) {
	for _, v := range arr {
		sum += v
	}

	return sum
}

func main() {
	var n int

	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}

	var arr []int
	var x int
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&x)
		if err != nil {
			fmt.Println(err)
			return
		}

		arr = append(arr, x)
	}

	fmt.Println(sum(arr))
}
