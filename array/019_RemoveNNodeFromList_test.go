package array

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	head := new(ListNode)
	head.Val = 1
	head.Next = new(ListNode)
	head.Next.Val = 2
	head.Next.Next = new(ListNode)
	head.Next.Next.Val = 3
	head.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Val = 4
	head.Next.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Next.Val = 5

	want := []int{1, 2, 3, 5}
	if got := removeNthFromEnd(head, 2); !isAMatch(want, got) {
		t.Errorf("removeNthFromEnd() = %v, want %v", got, want)
	}
}

func isAMatch(want []int, got *ListNode) bool {
	for _, element := range want {
		if got.Val != element {
			return false
		}
		got = got.Next
	}
	return true
}
