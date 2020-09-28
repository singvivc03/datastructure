package array

import (
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	want := 82
	target := 82
	nums := []int{1, 2, 4, 8, 16, 32, 64, 128}
	if got := threeSumClosest(nums, target); got != want {
		t.Errorf("threeSumClosest(%v, %d) = %d, want %d", nums, target, got, want)
	}

	want = -2
	target = -2
	nums = []int{4, 0, 5, -5, 3, 3, 0, -4, -5}
	if got := threeSumClosest(nums, target); got != want {
		t.Errorf("threeSumClosest(%v, %d) = %d, want %d", nums, target, got, want)
	}

	want = 1
	target = 1
	nums = []int{1, -3, 3, 5, 4, 1}
	if got := threeSumClosest(nums, target); got != want {
		t.Errorf("threeSumClosest(%v, %d) = %d, want %d", nums, target, got, want)
	}

	want = 3
	target = 0
	nums = []int{0, 1, 2}
	if got := threeSumClosest(nums, target); got != want {
		t.Errorf("threeSumClosest(%v, %d) = %d, want %d", nums, target, got, want)
	}

	want = -2
	target = -1
	nums = []int{-3, -2, -5, 3, -4}
	if got := threeSumClosest(nums, target); got != want {
		t.Errorf("threeSumClosest(%v, %d) = %d, want %d", nums, target, got, want)
	}

	want = 80
	target = 81
	nums = []int{1, 6, 9, 14, 16, 70}
	if got := threeSumClosest(nums, target); got != want {
		t.Errorf("threeSumClosest(%v, %d) = %d, want %d", nums, target, got, want)
	}

	want = 2
	target = 1
	nums = []int{-1, 2, 1, -4}
	if got := threeSumClosest(nums, target); got != want {
		t.Errorf("threeSumClosest(%v, %d) = %d, want %d", nums, target, got, want)
	}

	want = 2
	target = -100
	nums = []int{1, 1, 1, 0}
	if got := threeSumClosest(nums, target); got != want {
		t.Errorf("threeSumClosest(%v, %d) = %d, want %d", nums, target, got, want)
	}
}
