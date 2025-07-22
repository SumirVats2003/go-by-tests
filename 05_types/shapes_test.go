package types

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("calculating perimeter of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		expect := 40.0

		if got != expect {
			t.Errorf("got %.2f but expected %.2f", got, expect)
		}
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name   string
		shape  Shape
		expect float64
	}{
		{name: "Rectangle", shape: Rectangle{Height: 10, Width: 10}, expect: 100.0},
		{name: "Circle", shape: Circle{Radius: 10}, expect: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Height: 12, Base: 6}, expect: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.expect {
				t.Errorf("%#v got %g but expected %g", tt.shape, got, tt.expect)
			}
		})
	}
}
