package dp

import (
	"sort"
)

type order struct {
	start  int
	end    int
	volume int
}

func new(start, end, volume int) order {
	return order{start: start, end: start + end, volume: volume}
}

// By func
type By func(o1, o2 *order) bool

// Sort func
func (by By) Sort(orders []order) {
	or := &orderSorter{
		orders: orders,
		by:     by,
	}
	sort.Sort(or)
}

// Len is part of sort.Interface.
func (os *orderSorter) Len() int {
	return len(os.orders)
}

// Swap is part of sort.Interface.
func (os *orderSorter) Swap(i, j int) {
	os.orders[i], os.orders[j] = os.orders[j], os.orders[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (os *orderSorter) Less(i, j int) bool {
	return os.by(&os.orders[i], &os.orders[j])
}

type orderSorter struct {
	orders []order
	by     func(o1, o2 *order) bool
}

func phoneCalls(start []int, end []int, volumne []int) int {
	length := len(volumne)
	orders := make([]order, 0)
	for idx, v := range volumne {
		orders = append(orders, new(start[idx], end[idx], v))
	}
	increasingEndFunc := func(o1, o2 *order) bool {
		return o1.end < o2.end
	}
	By(increasingEndFunc).Sort(orders)
	return solve(orders, length)
}

func solve(orders []order, length int) int {
	dp := make([]int, length)
	dp[0] = orders[0].volume
	for i := 1; i < len(orders); i++ {
		currentVolume := orders[i].volume
		lastNonOverlappingIdx := binarySearch(orders, 0, i-1, orders[i].start)
		if lastNonOverlappingIdx != -1 {
			currentVolume += orders[lastNonOverlappingIdx].volume
		}
		dp[i] = max(currentVolume, dp[i-1])
	}
	return dp[length-1]
}

func binarySearch(orders []order, l int, r int, start int) int {
	for l < r-1 {
		mid := l + (r-1)/2
		if orders[mid].end <= start {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if orders[r].end < start {
		return r
	}
	if orders[l].end < start {
		return l
	}
	return -1
}
