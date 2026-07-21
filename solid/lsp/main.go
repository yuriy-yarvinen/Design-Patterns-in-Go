package main

import "fmt"

// Liskov Subtitution Principle

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}
func (r *Rectangle) SetWidth(width int) {
	r.width = width
}
func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

// wrong
// func (s *Square) SetHeight(height int) {
// 	s.height = height
// 	s.width = height
// }
// func (s *Square) SetWidth(height int) {
// 	s.height = height
// 	s.width = height
// }

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()

	fmt.Print("expected area ", expectedArea, " actual area is ", actualArea, "\n\n")
}

func main() {

	rc := &Rectangle{2, 3}
	UseIt(rc)

	sq := NewSquare(12)
	UseIt(sq)

}
