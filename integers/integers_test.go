package integers

import (
	"fmt"
	"testing"
)

//Please note that the example function will not be executed if you remove the comment // Output: 6.
//Although the function will be compiled, it won't be executed.
//By adding this code the example will appear in the documentation inside godoc, making your code even more accessible.
//To try this out, run godoc -http=:6060 and navigate to http://localhost:6060/pkg/
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		//we're using %d as our format strings rather than %q.
		//That's because we want it to print an integer rather than a string.
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

//To run the benchmarks do go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}
