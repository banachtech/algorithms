package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Append n copies of x to slice y
func appendn(y []int, x, n int) []int {
	for i := 0; i < n; i++ {
		y = append(y, x)
	}
	return y
}

// Histogram of slice values
func hist(y []int) map[int]int {
	h := make(map[int]int)
	for _, v := range y {
		h[v]++
	}
	return h
}

// Unfurl a slice of unique elements as per histogram
func unfurl(y []int, h map[int]int) []int {
	z := make([]int, 0)
	for _, v := range y {
		z = appendn(z, v, h[v])
	}
	return z
}

// Quicksort
func qsort2(y []int) []int {
	if len(y) <= 1 {
		return y
	}
	h := make(map[int]bool)
	for _, v := range y {
		if !h[v] {
			h[v] = true
		}
	}
	m := y[0]
	small, large := make([]int, 0), make([]int, 0)
	for k := range h {
		if k <= m {
			small = append(small, k)
		}
		if k > m {
			large = append(large, k)
		}
	}
	small = qsort2(small)
	large = qsort2(large)

	if large != nil {
		for _, v := range large {
			small = append(small, v)
		}
	}
	return small
}

func qsort(y []int) []int {
	h := hist(y)
	return unfurl(qsort2(y), h)
}

func pop(z *[]string) int {
	tmp, _ := strconv.Atoi((*z)[0])
	*z = (*z)[1:]
	return tmp
}

func main() {
	/* var n int
	// Parse input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	str := make([]string, 0)
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	n = pop(&str)
	y := make([]int, n)
	for i := range y {
		y[i] = pop(&str)
	}
	tmp := qSort3(y)
	for _, v := range tmp {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n") */

	var tq, tr, tt int64
	var st time.Time
	n := 10000
	for k := 0; k < 100; k++ {
		x := make([]int, n)
		for j := 0; j < n; j++ {
			x[j] = rand.Intn(10)
		}
		st = time.Now()
		_ = quickSort(x)
		tq += time.Since(st).Microseconds()
		st = time.Now()
		_ = qsort(x)
		tr += time.Since(st).Microseconds()
		st = time.Now()
		_ = qSort3(x)
		tt += time.Since(st).Microseconds()
	}
	fmt.Printf("method\ttime(microsec)\n")
	fmt.Printf("Quick\t%v\n", float64(tq/100))
	fmt.Printf("Quick2\t%v\n", float64(tr/100))
	fmt.Printf("Quick3\t%v\n", float64(tt/100))

}

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

// Quicksort
func qSort3(x []int) []int {
	if len(x) <= 1 {
		return x
	}
	m := x[0]
	eq := 0
	for i, v := range x {
		if v == m {
			x[eq], x[i] = x[i], x[eq]
			eq++
		}
	}
	lq := eq
	for i, v := range x {
		if v < m {
			x[lq], x[i] = x[i], x[lq]
			lq++
		}
	}
	gq := lq
	for i, v := range x {
		if v > m {
			x[gq], x[i] = x[i], x[gq]
			gq++
		}
	}
	s := quickSort(x[eq:lq])
	l := quickSort(x[lq:gq])

	for _, v := range x[:eq] {
		s = append(s, v)
	}
	for _, v := range l {
		s = append(s, v)
	}
	return s
}
