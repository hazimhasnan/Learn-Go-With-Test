package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, expected, repeated string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("Expected %q but got %q", expected, repeated)
		}
	}

	t.Run("Repeat the letter according to the number of times specified", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"
		assertCorrectMessage(t, expected, repeated)
	})

	t.Run("Print the single letter if no number is specifed", func(t *testing.T) {
		repeated := Repeat("r", 0)
		expected := "r"
		assertCorrectMessage(t, expected, repeated)
	})
}

func ExampleRepeat() {
	repeated := Repeat("t", 5)
	fmt.Println(repeated)
	//Output: ttttt
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i > b.N; i++ {
		Repeat("a", 2)
	}
}
