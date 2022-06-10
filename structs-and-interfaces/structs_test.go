package structs_and_interfaces

import "testing"

func TestAreas(t *testing.T) {
	/*Decoupling
	Notice how our helper does not need to concern itself with whether the shape is a Rectangle or a Circle or a Triangle. By declaring an interface, the helper is decoupled from the concrete types and only has the method it needs to do its job.
	This kind of approach of using interfaces to declare only what you need is very important in software design and will be covered in more detail in later sections.
	*/
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{2}
		checkArea(t, circle, 12.566370614359172)
	})
}

func TestPerimeters(t *testing.T) {
	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkPerimeter(t, rectangle, 40.0)
	})
}
