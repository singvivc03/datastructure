package dp

import (
	"testing"
)

func TestKnapsackRecur(t *testing.T) {
	want := 90
	weight, values := []int{5, 4, 6, 3}, []int{10, 40, 30, 50}
	if got := knapsackRecur(weight, values, 10, len(weight)); want != got {
		t.Errorf("knapsackRecur() = %d, wanted %d", got, want)
	}
}

func TestKnapsackDP(t *testing.T) {
	want := 90
	weight, values := []int{5, 4, 6, 3}, []int{10, 40, 30, 50}
	if got := knapsackDp(weight, values, 10, len(weight)); want != got {
		t.Errorf("knapsackDp() = %d, wanted %d", got, want)
	}
}
