package main

import (
	"fmt"
	"math/rand"
	"time"
)

func merge(a, b []int) []int {
	n, m := len(a), len(b)
	c := make([]int, n+m)
	var i, j int
	for k := range c {
		if i < n && j < m {
			if a[i] < b[j] {
				c[k] = a[i]
				i++
			} else {
				c[k] = b[j]
				j++
			}
		} else if i >= n && j < m {
			c[k] = b[j]
			j++
		} else if i < n && j >= m {
			c[k] = a[i]
			i++
		}
	}
	return c
}

// Mergesort
func mergeSort(x []int) []int {
	n := len(x)
	if n == 1 {
		return x
	}
	left := mergeSort(x[:n/2])
	right := mergeSort(x[n/2:])
	return merge(left, right)
}

// Quicksort
func quickSort(x []int) []int {
	if len(x) <= 1 {
		return x
	}
	m := x[0]
	small, large := make([]int, 0), make([]int, 0)
	for _, v := range x[1:] {
		if v <= m {
			small = append(small, v)
		}
		if v > m {
			large = append(large, v)
		}
	}
	small = quickSort(small)
	large = quickSort(large)

	small = append(small, m)
	if large != nil {
		for _, v := range large {
			small = append(small, v)
		}
	}
	return small
}

// Quicksort with random seed
func randSort(x []int) []int {
	if len(x) <= 1 {
		return x
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	k := r.Intn(len(x))
	x[0], x[k] = x[k], x[0]
	m := x[0]
	small, large := make([]int, 0), make([]int, 0)
	for _, v := range x[1:] {
		if v <= m {
			small = append(small, v)
		}
		if v > m {
			large = append(large, v)
		}
	}
	small = randSort(small)
	large = randSort(large)

	small = append(small, m)
	if large != nil {
		for _, v := range large {
			small = append(small, v)
		}
	}
	return small
}

// Tree sort
type tree struct {
	value       int
	left, right *tree
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func treeSort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func main() {
	var tm, tq, tr, tt int64
	var st time.Time
	n := 1000
	for k := 0; k < 100; k++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		x := make([]int, n)
		for j := 0; j < n; j++ {
			x[j] = r.Int()
		}
		st = time.Now()
		_ = mergeSort(x)
		tm += time.Since(st).Microseconds()
		st = time.Now()
		_ = quickSort(x)
		tq += time.Since(st).Microseconds()
		st = time.Now()
		_ = randSort(x)
		tr += time.Since(st).Microseconds()
		st = time.Now()
		treeSort(x)
		tt += time.Since(st).Microseconds()
	}
	fmt.Printf("method\ttime(microsec)\n")
	fmt.Printf("Merge\t%v\n", float64(tm/100))
	fmt.Printf("Quick\t%v\n", float64(tq/100))
	fmt.Printf("Rand\t%v\n", float64(tr/100))
	fmt.Printf("Tree\t%v\n", float64(tt/100))

}
