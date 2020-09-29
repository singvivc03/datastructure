package array

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeAllSortedList(lists)
}

func mergeAllSortedList(lists []*ListNode) *ListNode {
	n := len(lists)
	ans := newDummyListNode()
	temp := ans
	for true {
		min, minIdex := 1<<32, -1
		for j := 0; j < n; j++ {
			if lists[j] != nil {
				current := lists[j].Val
				if current < min {
					min = current
					minIdex = j
				}
			}
		}
		if min == 1<<32 {
			break
		}
		temp.Next = new(ListNode)
		temp.Next.Val = min
		temp = temp.Next
		lists[minIdex] = lists[minIdex].Next
	}
	return ans.Next
}

func newDummyListNode() *ListNode {
	dummy := new(ListNode)
	dummy.Val = -1
	return dummy
}
