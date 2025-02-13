/*
 * (10 points)
 * Define an interface Shape to represent a 2D shape. The interface
 * should contain:
 * - a function that computes the area of the shape,
 * - a function that returns the position of the shape,
 * - a function that moves the position of the shape.
 * Define a struct for some 2D shape and make the struct implement the
 * Shape interface. Also define a function Test that accepts a Shape
 * as its input parameter. Print the position and the area of the shape
 * in the function Test. Run Test with some concrete shape in the main
 * function.
 *
 * Hint:
 * - You may define another struct for positions.
 */

// -- <code begin> --
package main

import "fmt"

type position struct {
	x, y int
}

type Shape interface {
	area() int
	position() position
	moves(x int, y int)
}

type Rectangle struct {
	width, height int
	pos           position
}

func (s Rectangle) area() int {
	return s.width * s.height
}

func (s Rectangle) position() position {
	return s.pos
}

func (s Rectangle) moves(x int, y int) {
	s.pos.x = x
	s.pos.y = y
}
func Test(s Shape) {
	fmt.Printf("位置: (%d, %d)\n", s.position().x, s.position().y)
	fmt.Printf("面積: %d\n", s.area())
}

func main() {
	rt := Rectangle{
		width:  8,
		height: 6,
		pos:    position{-4, 3},
	}

	Test(rt)
}

// -- <code end> --
