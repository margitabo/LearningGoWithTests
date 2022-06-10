package structs_and_interfaces

import "testing"

func TestAreas(t *testing.T) {
	t.Run("rectanglesArea", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circleArea", func(t *testing.T) {
		circle := Circle{2}
		got := circle.Area()
		want := 12.56

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestPerimeters(t *testing.T) {
	t.Run("rectanglesPerimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
