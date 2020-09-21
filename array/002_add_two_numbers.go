package array

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := new(ListNode)
	ans.Val = -1
	temp := ans
	carry := 0
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + carry
		carry = val / 10
		val %= 10
		temp.Next = new(ListNode)
		temp.Next.Val = val
		temp = temp.Next
		l1, l2 = l1.Next, l2.Next
	}
	temp, carry = checkLeftOver(l1, temp, &carry)
	temp, _ = checkLeftOver(l2, temp, &carry)
	if carry > 0 {
		temp.Next = new(ListNode)
		temp.Next.Val = carry
	}
	return ans.Next
}

func checkLeftOver(ln *ListNode, temp *ListNode, carry *int) (*ListNode, int) {
	for ln != nil {
		val := ln.Val + *carry
		*carry = val / 10
		val %= 10
		temp.Next = new(ListNode)
		temp.Next.Val = val
		temp = temp.Next
		ln = ln.Next
	}
	return temp, *carry
}
