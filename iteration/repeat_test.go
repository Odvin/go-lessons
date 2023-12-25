package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 'a' five times", func(t *testing.T) {
		assertCorrectResponse(t, "aaaaa", Repeat("a", 5))
	})

	t.Run("repeat 'b' three times", func(t *testing.T) {
		assertCorrectResponse(t, "bbb", Repeat("b", 3))
	})

}

func assertCorrectResponse(t testing.TB, expected, got string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

func ExampleRepeat() {
	res := Repeat("x", 3)
	fmt.Println(res)
	// Output: xxx
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
