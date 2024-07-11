package main

import "fmt"

type Rectangle struct {
	h int
	w int
}

func (rect *Rectangle) getArea() int {
	return rect.h * rect.w
}

func main() {
	rect := Rectangle{5, 10}
	fmt.Println(rect.getArea())
}
