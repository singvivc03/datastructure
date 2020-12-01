package dp

import (
	"testing"
)

func TestPhoneCalls(t *testing.T) {
	want := 4
	start, end, volume := []int{1, 2, 4}, []int{2, 2, 1}, []int{1, 3, 3}
	result := phoneCalls(start, end, volume)
	if result != want {
		t.Errorf("phoneCalls(%v, %v, %v)=%d, want %d", start, end, volume, result, want)
	}
}
