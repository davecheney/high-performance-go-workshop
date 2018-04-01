package fib

// STARTBENCH OMIT
import "testing"

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(20) // run the Fib function b.N times
	}
}

// ENDBENCH OMIT

// STARTFIB OMIT
// Fib computes the n'th number in the Fibonacci series.
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

// ENDFIB OMIT

func Fib2(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

func TestFib(t *testing.T) {
	fibs := []int{0, 1, 1, 2, 3, 5, 8, 13, 21}
	for n, want := range fibs {
		got := Fib(n)
		if want != got {
			t.Errorf("Fib(%d): want %d, got %d", n, want, got)
		}
	}
}

func TestFib2(t *testing.T) {
	fibs := []int{0, 1, 1, 2, 3, 5, 8, 13, 21}
	for n, want := range fibs {
		got := Fib2(n)
		if want != got {
			t.Errorf("Fib2(%d): want %d, got %d", n, want, got)
		}
	}
}

func TestFibFib(t *testing.T) {
	for n := 0; n < 30; n++ {
		want := Fib(n)
		got := Fib2(n)
		if want != got {
			t.Errorf("Fib2(%d): want %d, got %d", n, want, got)
		}
	}
}

func _BenchmarkFib2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib2(20)
	}
}
