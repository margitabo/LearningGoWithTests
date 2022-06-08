package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			//It is sometimes useful to also print the inputs to the function in the error message.
			//Here, we are using the %v placeholder to print the "default" format, which works well for arrays
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4})

	want := []int{3, 7}
	//It's important to note that reflect.DeepEqual is not "type safe" -
	//the code will compile even if you did something a bit silly.
	//want := "blob"

	//Go does not let you use equality operators with slices
	//invalid operation: got != want (slice can only be compared to nil)
	if !reflect.DeepEqual(got, want) {
		//It is sometimes useful to also print the inputs to the function in the error message.
		//Here, we are using the %v placeholder to print the "default" format, which works well for arrays
		t.Errorf("got %d want %d given", got, want)
	}
}
