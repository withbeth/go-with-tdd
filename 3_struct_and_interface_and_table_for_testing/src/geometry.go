package main

import "math"

// Requirement:
// We want to be able to write some kind of "checkArea" function
// that we can pass both of "Rectangle" and "Circle" to,
// but fail to compile if we try and pass in something that isn't a shape.
// In Go, we can codify the intent with "interfaces"
// Interfaces are a very powerful concept in statically typed lang like Go.
// They allow you to make functions that can be used with different types
// and create decoupled code while still maintaining type-safety.

// This is quite different to interfaces in most other programming langs.
// Other lang : MyType Foo implements Interface Bar
// Go : Interface solution is implicit. if the type you pass in matches
// what the interface is asking for, it will compile.
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}
// Declare methods
// Convention : Have a receiver var be the first letter of the type.
func (r Rectangle) Area() (float64) {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}
// Declare method
func (c Circle) Area() (float64) {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base float64
	Height float64
}

func (t Triangle) Area() (float64) {
	return (t.Base * t.Height) / 2
}

// Calculate a perimeter of a rectangle given a height and width
func Perimeter(rectangle Rectangle)(float64){
	return 2 * (rectangle.Width + rectangle.Height)
}

// In Go lang, you can not declare same func with diff param types..
// Opt 1) Have diff package.
// Opt 2) define "methods"
// Methods are very similar to "Functions". but they are called by invoking them
// on an instance of a particular type.
// Calculate a area of a rectangle given a height and width
func Area(rectangle Rectangle)(float64) {
	return rectangle.Width * rectangle.Height
}



