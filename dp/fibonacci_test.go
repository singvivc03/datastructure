package dp

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	n, want := 5, int64(5)
	if got := Fibonacci(n); got != want {
		t.Errorf("Fibonacci(%d) = %d, want %d", n, got, want)
	}
	n, want = 6, int64(8)
	if got := Fibonacci(n); got != want {
		t.Errorf("Fibonacci(%d) = %d, want %d", n, got, want)
	}
}

func TestFibonacciTabular(t *testing.T) {
	n, want := 5, int64(5)
	if got := FibonacciTabular(n); got != want {
		t.Errorf("FibonacciTabular(%d) = %d, want %d", n, got, want)
	}
	n, want = 6, int64(8)
	if got := FibonacciTabular(n); got != want {
		t.Errorf("FibonacciTabular(%d) = %d, want %d", n, got, want)
	}
}

func TestFibonacciWithConstantSpace(t *testing.T) {
	n, want := 5, int64(5)
	if got := FibonacciWithConstantSpace(n); got != want {
		t.Errorf("FibonacciWithConstantSpace(%d) = %d, want %d", n, got, want)
	}
	n, want = 6, int64(8)
	if got := FibonacciWithConstantSpace(n); got != want {
		t.Errorf("FibonacciWithConstantSpace(%d) = %d, want %d", n, got, want)
	}
}
