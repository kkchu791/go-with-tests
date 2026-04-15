package perimeter2

import (
	"math"
	"testing"
)

type Shape interface {
	CalculateArea() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) CalculateArea() float64 {
	return (r.Width * r.Height)
}

func (c Circle) CalculateArea() float64 {
	return math.Pow(c.Radius, 2) * 3.14
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := float64(40.0)

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.CalculateArea()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{2, 2}
		checkArea(t, rectangle, 4)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314)
	})
}

func Perimeter(r Rectangle) float64 {
	return float64(r.Width+r.Height) * 2
}

func CalculateArea(r Rectangle) float64 {
	return float64(r.Width * r.Height)
}
