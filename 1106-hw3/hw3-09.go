package main

import "fmt"

/*
 * (10 points)
 * Complete the function printArea such that it prints the
 * area of the input shape, which may be one of circle,
 * square, or rectangle. You may use type assertion or
 * type switch.
 *
 * The expected output is:
 *   Area of circle   : 314.159265359
 *   Area of square   : 6.25
 *   Area of rectangle: 28
 */

var pi float64 = 3.14159265359

// area = radius * radius * pi
type circle struct {
	radius float64
}

// area = width * width
type square struct {
	width float64
}

// area = width * height
type rectangle struct {
	width  float64
	height float64
}

func printArea(shape interface{}) {
	// -- <code begin> --

	switch s := shape.(type) {
	case circle:
		a := s.radius * s.radius * pi
		fmt.Printf("Area of circle   : %.9f\n", a)
	case square:
		a := s.width * s.width
		fmt.Printf("Area of square   : %.2f\n", a)
	case rectangle:
		a := s.width * s.height
		fmt.Printf("Area of rectangle: %.0f\n", a)
	}
	// -- <code end> --
}

func main() {
	printArea(circle{10})
	printArea(square{2.5})
	printArea(rectangle{4, 7})
}
