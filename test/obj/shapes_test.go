package obj

import (
	"math"
	"testing"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

func TestPerimeter(t *testing.T) {
	t.Run("rectangle perimeter", func(t *testing.T) {
		got := perimeter(Rectangle{10.0, 10.0})
		want := 40.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		area  float64
	}{
		{name: "Rectangle", shape: Rectangle{Height: 10.0, Width: 10.0}, area: 100},
		{name: "Circle", shape: Circle{Radius: 10}, area: math.Pi * 10 * 10},
		{name: "Triangle", shape: Triangle{Base: 5, Height: 10}, area: 25},
	}

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("%#v got %.2f want %.2f", shape, got, want)
		}
	}

	for _, areaTest := range areaTests {
		t.Run(areaTest.name, func(t *testing.T) {
			checkArea(t, areaTest.shape, areaTest.area)
		})
	}
}

func perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2
}
