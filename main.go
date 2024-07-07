package main

import "fmt"

type Rectangle struct {
	h int
	w int
}

func getArea(rect Rectangle) int {
	return rect.h * rect.w
}

func main() {
	rect := Rectangle{5, 10}
	fmt.Println(getArea(rect))
}
