package array

import (
	"testing"
)

func TestSwapPairs(t *testing.T) {
	want := []int{2, 1, 4, 3}
	head := createList([]int{1, 2, 3, 4})
	if got := swapPairs(head); !isAMatch(want, got) {
		t.Errorf("swapPairs(%v)=%v, want %v", head, got, want)
	}
}

func createList(list []int) *ListNode {
	dummy := new(ListNode)
	dummy.Val = -1
	tmp := dummy
	for i := 0; i < len(list); i++ {
		tmp.Next = new(ListNode)
		tmp.Next.Val = list[i]
		tmp = tmp.Next
	}
	return dummy.Next
}
