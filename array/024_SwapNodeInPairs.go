package array

func swapPairs(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Val, dummy.Next = 0, head
	current := dummy

	for current.Next != nil && current.Next.Next != nil {
		firstNode := current.Next
		secondNode := current.Next.Next
		firstNode.Next = secondNode.Next
		current.Next = secondNode
		current.Next.Next = firstNode
		current = current.Next.Next
	}
	return dummy.Next
}
