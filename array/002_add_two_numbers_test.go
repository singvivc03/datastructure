package array

import "testing"

// TestAddTwoNumbers ...
func TestAddTwoNumbers(t *testing.T) {
	l1, l2 := new(ListNode), new(ListNode)
	l1.Val = 2
	l1.Next = new(ListNode)
	l1.Next.Val = 4
	l1.Next.Next = new(ListNode)
	l1.Next.Next.Val = 3

	l2.Val = 5
	l2.Next = new(ListNode)
	l2.Next.Val = 6
	l2.Next.Next = new(ListNode)
	l2.Next.Next.Val = 4

	want := new(ListNode)
	want.Val = 7
	want.Next = new(ListNode)
	want.Next.Val = 0
	want.Next.Next = new(ListNode)
	want.Next.Next.Val = 8

	if got := addTwoNumbers(l1, l2); !checkEqual(want, got) {
		t.Errorf("addTwoNumbers(%v, %v) = %v, want %v", l1, l2, got, want)
	}

	l1, l2 = new(ListNode), new(ListNode)
	l1.Val = 1
	l1.Next = new(ListNode)
	l1.Next.Val = 8

	l2.Val = 0

	want = new(ListNode)
	want.Val = 1
	want.Next = new(ListNode)
	want.Next.Val = 8

	if got := addTwoNumbers(l1, l2); !checkEqual(want, got) {
		t.Errorf("addTwoNumbers(%v, %v) = %v, want %v", l1, l2, got, want)
	}

	l1, l2 = new(ListNode), new(ListNode)
	l1.Val = 8
	l1.Next = new(ListNode)
	l1.Next.Val = 6

	l2.Val = 6
	l2.Next = new(ListNode)
	l2.Next.Val = 4
	l2.Next.Next = new(ListNode)
	l2.Next.Next.Val = 8

	want = new(ListNode)
	want.Val = 4
	want.Next = new(ListNode)
	want.Next.Val = 1
	want.Next.Next = new(ListNode)
	want.Next.Next.Val = 9

	if got := addTwoNumbers(l1, l2); !checkEqual(want, got) {
		t.Errorf("addTwoNumbers(%v, %v) = %v, want %v", l1, l2, got, want)
	}
}

func checkEqual(want *ListNode, got *ListNode) bool {
	for want != nil {
		if want.Val != got.Val {
			return false
		}
		want, got = want.Next, got.Next
	}
	return true
}
