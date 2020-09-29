package array

import (
	"testing"
)

func TestMergeKSortedList(t *testing.T) {
	want := []int{1, 1, 2, 3, 4, 4, 5, 6}
	lists := getInputLists([][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}})
	if got := mergeKLists(lists); !isResultMatching(want, got) {
		t.Errorf("mergeKLists(%v)=%v, want %v", lists, got, want)
	}

	want = []int{1, 2, 3}
	lists = getInputLists([][]int{{1, 2}, {3}})
	if got := mergeKLists(lists); !isResultMatching(want, got) {
		t.Errorf("mergeKLists(%v)=%v, want %v", lists, got, want)
	}

	want = []int{1, 2, 3}
	lists = getInputLists([][]int{{1}, {3}, {2}})
	if got := mergeKLists(lists); !isResultMatching(want, got) {
		t.Errorf("mergeKLists(%v)=%v, want %v", lists, got, want)
	}
}

func isResultMatching(expected []int, actual *ListNode) bool {
	temp := actual
	for i := 0; i < len(expected); i++ {
		if temp.Val != expected[i] {
			return false
		}
		temp = temp.Next
	}
	return true
}

func getInputLists(input [][]int) []*ListNode {
	var lists []*ListNode
	for i := 0; i < len(input); i++ {
		list := new(ListNode)
		list.Val = -1
		temp := list
		for j := 0; j < len(input[i]); j++ {
			temp.Next = new(ListNode)
			temp.Next.Val = input[i][j]
			temp = temp.Next
		}
		lists = append(lists, list.Next)
	}
	return lists
}
