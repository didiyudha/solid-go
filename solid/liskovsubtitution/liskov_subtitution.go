package main

import "fmt"

type Shape interface {
	GetHeight() int
	GetWidth() int
}

type Square struct {
	Height int
	Width  int
}

type Rectangle struct {
	Height int
	Width  int
}

func (s *Square) GetHeight() int {
	return s.Height
}

func (s *Square) GetWidth() int {
	return s.Width
}

func (r *Rectangle) GetHeight() int {
	return r.Height
}

func (r *Rectangle) GetWidth() int {
	return r.Width
}

func Size(shape Shape) int {
	h := shape.GetHeight()
	w := shape.GetWidth()
	return h * w
}

func main() {
	rect := Rectangle{Height: 10, Width: 5}
	square := Square{Height: 10, Width: 10}

	fmt.Println("Rectangle size: ", Size(&rect))
	fmt.Println("Square size: ", Size(&square))
}
