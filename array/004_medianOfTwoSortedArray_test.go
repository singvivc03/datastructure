package array

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	want := 5.5
	nums1 := []int{2, 5, 6, 11}
	nums2 := []int{3, 4, 9, 10}
	if got := findMedianSortedArrays(nums1, nums2); got != want {
		t.Errorf("findMedianSortedArrays(%v, %v)=%f, want %f", nums1, nums2, got, want)
	}
	want = 2.5
	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	if got := findMedianSortedArrays(nums1, nums2); got != want {
		t.Errorf("findMedianSortedArrays(%v, %v)=%f, want %f", nums1, nums2, got, want)
	}

	want = 13.5
	nums1 = []int{23, 26, 31, 35}
	nums2 = []int{3, 5, 7, 9, 11, 16}
	if got := findMedianSortedArrays(nums1, nums2); got != want {
		t.Errorf("findMedianSortedArrays(%v, %v)=%f, want %f", nums1, nums2, got, want)
	}

	want = 2.0
	nums1 = []int{1, 3}
	nums2 = []int{2}
	if got := findMedianSortedArrays(nums1, nums2); got != want {
		t.Errorf("findMedianSortedArrays(%v, %v)=%f, want %f", nums1, nums2, got, want)
	}

	want = 2.5
	nums1 = []int{}
	nums2 = []int{2, 3}
	if got := findMedianSortedArrays(nums1, nums2); got != want {
		t.Errorf("findMedianSortedArrays(%v, %v)=%f, want %f", nums1, nums2, got, want)
	}

	want = 2.5
	nums1 = []int{2, 3}
	nums2 = []int{}
	if got := findMedianSortedArrays(nums1, nums2); got != want {
		t.Errorf("findMedianSortedArrays(%v, %v)=%f, want %f", nums1, nums2, got, want)
	}

	want = 0.0
	nums1 = []int{}
	nums2 = []int{}
	if got := findMedianSortedArrays(nums1, nums2); got != want {
		t.Errorf("findMedianSortedArrays(%v, %v)=%f, want %f", nums1, nums2, got, want)
	}
}
