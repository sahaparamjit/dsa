package main

import (
	"fmt"
	"math"
)

// Point cooridinate
type Point struct {
	X int
	Y int
}

// Square struct
type Square struct {
	Point Point
	Length int
}

func (sq *Square) distanceFromOrigin() float64{
	x := sq.Point.X
	y := sq.Point.X
	return math.Sqrt((math.Pow(float64(x), 2) + math.Pow(float64(y), 2)))
}

func (p *Point) move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

func (sq *Square) area() float64{
	return math.Pow(float64(sq.Length), 2)
}

// NewSquare can create a new square with specified args
func NewSquare(x int, y int, length int) (*Square, error){
	if x<0 || y<0 {
		return nil, fmt.Errorf("Center coordinate cannot be ")
	} else if length <= 0 {
		return nil, fmt.Errorf("Length should be greater than 0")
	}

	return &Square{Point: Point{x, y}, Length: length}, nil
}

func main() {
	v, err := NewSquare(0, 0, 12)
	if err != nil {
		fmt.Printf("Error caused due to %v", err)
	}
	fmt.Printf("%+v\n", v.area())
	v.Point.move(4 , 3)
	fmt.Printf("%+v\n", v.distanceFromOrigin())
}