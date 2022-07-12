package main

import (
	"fmt"
)

type Heap []int

func NewHeap(x []int) Heap {
	var h Heap
	for i := range x {
		h.Insert(x[i])
	}
	return h
}

func (h *Heap) Insert(x int) {
	*h = append(*h, x)
	n := len(*h) - 1
	for (*h)[n] < (*h)[n>>1] && n > 0 {
		(*h)[n], (*h)[n>>1] = (*h)[n>>1], (*h)[n]
		n = n >> 1
	}
}

func (h *Heap) ExtractMin() int {
	min := (*h)[0]
	n := len(*h) - 1
	(*h)[0] = (*h)[n]
	*h = (*h)[:n]
	if len(*h) <= 1 {
		return min
	}
	i := 0
	for i < n-1 {
		j := i<<1 + 1
		if j == n-1 {
			if (*h)[i] > (*h)[j] {
				(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
				break
			}
		}
		if j <= n-2 {
			if (*h)[j] < (*h)[j+1] {
				(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
				i = j
			} else {
				(*h)[i], (*h)[j+1] = (*h)[j+1], (*h)[i]
				i = j + 1
			}
		}
	}
	return min
}

func main() {
	x := []int{1, 3, 4, 0, 3, 2, 1}
	h := NewHeap(x)
	fmt.Println(h)
	fmt.Println(h.ExtractMin())
	fmt.Println(h)
}
