package dp

import (
	"testing"
)

func TestProbabilityWithMoreHeads(t *testing.T) {
	want := 0.612
	n := 3
	probabilities := []float64{0.30, 0.60, 0.80}

	if got := probabilityWithMoreHeads(n, probabilities); got != want {
		t.Errorf("probabilityWithMoreHeads(%d, %v) = %f, want %f", n, probabilities, got, want)
	}

	want = 0.5
	n = 1
	probabilities = []float64{0.50}

	if got := probabilityWithMoreHeads(n, probabilities); got != want {
		t.Errorf("probabilityWithMoreHeads(%d, %v) = %f, want %f", n, probabilities, got, want)
	}

	want = 0.3821815872
	n = 5
	probabilities = []float64{0.42, 0.01, 0.42, 0.99, 0.42}

	if got := probabilityWithMoreHeads(n, probabilities); got != want {
		t.Errorf("probabilityWithMoreHeads(%d, %v) = %f, want %f", n, probabilities, got, want)
	}
}
