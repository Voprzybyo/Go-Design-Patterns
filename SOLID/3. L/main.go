package main

import "fmt"

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

func NewBetterSquare(size int) *BetterSquare {
	sq := BetterSquare{size}
	return &sq
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

type BetterSquare struct {
	size int
}

func (s *BetterSquare) Rectangle() *Rectangle {
	return &Rectangle{s.size, s.size}
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Print("Expected an area of ", expectedArea, ", got ", actualArea, "\n")
}

func main() {

	rc := &Rectangle{2, 3}
	UseIt(rc)

	// Wrong approach
	sq := NewSquare(5)
	UseIt(sq)

	// Better solution -> Liskov substitution principle
	bettersq := NewBetterSquare(5)
	sqrc := bettersq.Rectangle()
	UseIt(sqrc)
}
