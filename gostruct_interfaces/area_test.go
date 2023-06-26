package gostruct_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.2, 12.3}
	got := Perimeter(rectangle)
	want := float64(45)

	if got != want {
		t.Errorf("got %.2f but want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{10.0, 20.0}
	got := rectangle.Area()
	want := float64(200)

	if got != want {
		t.Errorf("got %.2f but want %.2f", got, want)
	}
}

func TestShapes(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f but want %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 12.0}
		// got := rectangle.Area()
		// want := float64(144)
		checkArea(t, rectangle, 144.00)

	})

	// we can use %g instead of %f to get more precised value

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		// got := circle.Area()
		// want := 314.1592653589793
		checkArea(t, circle, 314.59)
	})
}

func TestArea1(t *testing.T) {
	// slice of structs
	areatest := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{length: 12.0, width: 6.0}, want: 72.0},
		{shape: Circle{radius: 10.0}, want: 314.1592653589793},
		{shape: Triangle{height: 12.0, base: 6.0}, want: 36.0},
	}

	for _, tt := range areatest {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g but want %g", got, tt.want)
		}
	}
}

//Table driven tests can be a great item in your toolbox, but be sure that you have a need for the extra noise in the tests.
//They are a great fit when you wish to test various implementations of an interface, or
//if the data being passed in to a function has lots of different requirements that need testing.

//Declaring structs to create your own data types which lets you bundle related data together and make the intent of your code clearer
//Declaring interfaces so you can define functions that can be used by different types (parametric polymorphism)
//Adding methods so you can add functionality to your data types and so you can implement interfaces
//Table driven tests to make your assertions clearer and your test suites easier to extend & maintain

func TestArea2(t *testing.T) {

	checkArea := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %g but want %g", got, want)
		}
	}
	// slice of structs
	areatest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{length: 12.0, width: 6.0}, want: 72.0},
		{name: "Circle", shape: Circle{radius: 10.0}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{height: 12.0, base: 6.0}, want: 36.0},
	}

	for _, tt := range areatest {
		got := tt.shape.Area()
		checkArea(t, got, tt.want)
	}
}

// Interfaces are a great tool for hiding complexity away from other parts of the system
