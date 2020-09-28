package array

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	ans := new(ListNode)
	ans.Val = -1
	temp := ans
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			temp.Next = new(ListNode)
			temp.Next.Val = l1.Val
			temp = temp.Next
			l1 = l1.Next
		} else {
			temp.Next = new(ListNode)
			temp.Next.Val = l2.Val
			temp = temp.Next
			l2 = l2.Next
		}
	}
	for l1 != nil {
		temp.Next = new(ListNode)
		temp.Next.Val = l1.Val
		temp, l1 = temp.Next, l1.Next
	}
	for l2 != nil {
		temp.Next = new(ListNode)
		temp.Next.Val = l2.Val
		temp, l2 = temp.Next, l2.Next
	}
	return ans.Next
}
