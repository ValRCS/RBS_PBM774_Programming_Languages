//let's talk about interfaces in Go
//interfaces in Go are a way to define behavior
//interfaces in Go are a way to define a set of methods
//interfaces in Go are a way to define a set of methods that a type must implement

//let's create an interface called Shape which will have Area and Perimeter methods
//we can create an interface in Go using the interface keyword

package main

import (
	"fmt"
)

type Shape interface {
	Area() float64 //so Area should return a float64
	Perimeter() float64 //Perimeter should return a float64
}
//so technically I could skip having an interface and just have the methods in the structures and it would still compile and work
//TODO explore the why behind interfaces in Go, and why they are useful
//TIP programming by contract
//TIP interfaces in Go are a way to define behavior

//let's creat a Rectangle structure

type Rectangle struct {
	Width  float64
	Height float64
}

//let's create a Circle structure
type Circle struct {
	Radius float64
}

//let's create a method for Rectangle called Area
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//let's create a method for Rectangle called Perimeter
func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

//let's create a method for Circle called Area
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

//let's create a method for Circle called Perimeter
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func main() {
	//let's create a Rectangle
	r := Rectangle{
		Width:  5,
		Height: 4,
	}

	//let's create a Circle
	c := Circle{
		Radius: 10,
	}

	//let's print the area of the Rectangle
	fmt.Println("Area of the Rectangle is", r.Area())

	//let's print the perimeter of the Rectangle
	fmt.Println("Perimeter of the Rectangle is", r.Perimeter())

	//let's print the area of the Circle
	fmt.Println("Area of the Circle is", c.Area())

	//let's print the perimeter of the Circle
	fmt.Println("Perimeter of the Circle is", c.Perimeter())	
}