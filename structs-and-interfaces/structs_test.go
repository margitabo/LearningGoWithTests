package structs_and_interfaces

import "testing"

/*
func TestAreas(t *testing.T) {
	//Decoupling
	//Notice how our helper does not need to concern itself with whether the shape is a Rectangle or a Circle or a Triangle. By declaring an interface, the helper is decoupled from the concrete types and only has the method it needs to do its job.
	//This kind of approach of using interfaces to declare only what you need is very important in software design and will be covered in more detail in later sections.

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
*/

//Table driven tests  are useful when you want to build a list of test cases that can be tested in the same manner.
//Table driven tests can be a great item in your toolbox, but be sure that you have a need for the extra noise in the tests.
//They are a great fit when you wish to test various implementations of an interface, or if the data being passed in to a function has lots of different requirements that need testing.
func TestArea(t *testing.T) {

	areaTests := []struct {
		name     string
		shape    Shape
		wantArea float64
	}{
		//{Rectangle{12, 6}, 72.0},
		//{Circle{10}, 3.141592653589793e+10},
		//{Triangle{12, 6}, 36.0},
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, wantArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, wantArea: 3.141592653589793e+10},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, wantArea: 36.0},
		//One final tip with table driven tests is to use t.Run and to name the test cases.
		//By wrapping each case in a t.Run you will have clearer test output on failures as it will print the name of the case
		//And you can run specific tests within your table with go test -run TestArea/Rectangle
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.wantArea {
				//We can change our error message into %#v got %g want %g. The %#v format string will print out our struct with the values in its field, so the developer can see at a glance the properties that are being tested.

				t.Errorf("%#v got %g want %g", tt.shape, got, tt.wantArea)
			}
		})
	}
}
