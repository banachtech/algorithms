// Compute max of products two integers in a sequence
// Brute force: compute pairwise product (of upper triangular) and pick max. O(n^2)
// Smart way: Pick two largest integers and compute product
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func max(m []int) (int, int) {
	out := 0
	j := 0
	for i, v := range m {
		if v > out {
			out = v
			j = i
		}
	}
	return j, out
}

func maxprod(m []int) int {
	n, x := max(m)
	m[0], m[n] = m[n], m[0]
	_, y := max(m[1:])
	return x * y
}

func main() {
	a, _ := strconv.Atoi(os.Args[1])
	m := make([]int, a)
	for i := range m {
		m[i] = rand.Intn(200000)
	}
	st := time.Now()
	out := maxprod(m)
	fmt.Printf("%d\t%v\n", out, time.Since(st))
}

// Even faster implementation of maxprod
// Sort input array and take the last two values!
