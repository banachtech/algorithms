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

func mergeSort(x []int) []int {
	n := len(x)
	if n == 1 {
		return x
	}
	left := mergeSort(x[:n/2])
	right := mergeSort(x[n/2:])
	return merge(left, right)
}

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

func main() {
	var tm, tq, tr int64
	var st time.Time
	n := 10000
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
	}
	fmt.Printf("method\ttime\n")
	fmt.Printf("Merge\t%v\n", float64(tm/100))
	fmt.Printf("Quic\t %v\n", float64(tq/100))
	fmt.Printf("Rand\t%v\n", float64(tr/100))

}
