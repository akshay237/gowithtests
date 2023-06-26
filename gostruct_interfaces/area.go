package gostruct_interfaces

import "math"

type Rectangle struct {
	length, width float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	height float64
	base   float64
}

// Methods are very similar to functions but they are called by invoking them on an instance of a particular type.
// When your method is called on a variable of that type, you get your reference to its data via the receiverName variable.

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.length + r.width)
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

// functions of interfaces are implemented by multipe structs if they satisfies function signature
// By declaring an interface, the helper is decoupled from the concrete types and only has the method it needs to do its job.
type Shape interface {
	Area() float64
}
