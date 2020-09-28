package array

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	want := []int{1, 1, 2, 3, 4, 4}
	l1, l2 := getList([]int{1, 2, 4}), getList([]int{1, 3, 4})

	if got := mergeTwoLists(l1.Next, l2.Next); !isAMatch(want, got) {
		t.Errorf("mergeTwoLists(%v, %v)=%v, want %v", l1.Next, l2.Next, got, want)
	}
}

func getList(nums []int) *ListNode {
	dummy := new(ListNode)
	dummy.Val = -1
	l1 := dummy
	for i := 0; i < len(nums); i++ {
		dummy.Next = new(ListNode)
		dummy.Next.Val = nums[i]
		dummy = dummy.Next
	}
	return l1
}
