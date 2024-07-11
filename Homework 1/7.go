package main

import (
	"fmt"
)

func getPrimes(n int) []int {
	isPrime := make([]bool, n+1)
	for i := 0; i < n+1; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	var res []int
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			res = append(res, i)
		}
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

	for _, v := range getPrimes(n) {
		fmt.Println(v)
	}
}
