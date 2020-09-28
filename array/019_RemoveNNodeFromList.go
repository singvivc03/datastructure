package array

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := new(ListNode)
	dummy.Val, dummy.Next = -1, head
	first, second := dummy, dummy
	for n >= 0 {
		first = first.Next
		n--
	}
	for first != nil {
		second = second.Next
		first = first.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}
