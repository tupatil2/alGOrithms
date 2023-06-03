package main

import (
	"fmt"
)

type SegmentTreeInterface interface {
	Build(index int, low int, high int, arr []int)
	Update(index int, low int, high int, i int, val int)
	FindSum(index int, low int, high int, l int, r int) int
}

type SegmentTree struct{}

var seg []int

func (s *SegmentTree) Build(index int, low int, high int, arr []int) {
	if low == high {
		seg[index] = arr[low]
		return
	}
	mid := low + (high-low)/2
	s.Build(2*index+1, low, mid, arr)
	s.Build(2*index+2, mid+1, high, arr)
	seg[index] = seg[2*index+1] + seg[2*index+2]
}

func (s *SegmentTree) Update(index int, low int, high int, i int, val int) {
	if low == high {
		seg[index] = val
		return
	}
	mid := low + (high-low)/2
	if i <= mid {
		s.Update(2*index+1, low, mid, i, val)
	} else {
		s.Update(2*index+2, mid+1, high, i, val)
	}
	seg[index] = seg[2*index+1] + seg[2*index+2]
}

func (s *SegmentTree) FindSum(index int, low int, high int, l int, r int) int {
	if r < low || l > high {
		return 0
	}
	if low >= l && high <= r {
		return seg[index]
	}

	mid := low + (high-low)/2
	left := s.FindSum(2*index+1, low, mid, l, r)
	right := s.FindSum(2*index+2, mid+1, high, l, r)
	return left + right
}

func main() {
	var segmentTree SegmentTreeInterface
	segmentTree = &SegmentTree{}

	n := 5
	arr := []int{1, 2, 3, 4, 5}
	seg = make([]int, 4*n+1)
	segmentTree.Build(0, 0, n-1, arr)

	// Find the sum in a given range
	left := 0
	right := n-1
	sum := segmentTree.FindSum(0, 0, n-1, left, right)
	fmt.Printf("Sum of elements from index %d to %d: %d\n", left, right, sum)

  segmentTree.Update(0, 0, n-1, 2, 10)
	sum = segmentTree.FindSum(0, 0, n-1, left, right)
	fmt.Printf("Sum of elements from index %d to %d: %d\n", left, right, sum)
}
