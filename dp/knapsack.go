package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func max(a ...int) int {
	tmp := a[0]
	for _, v := range a {
		if v > tmp {
			tmp = v
		}
	}
	return tmp
}

func pop(z *[]string) int {
	tmp, _ := strconv.Atoi((*z)[0])
	*z = (*z)[1:]
	return tmp
}

// compute max stuffed value and objects stuffed
func stuffem(V, S []int, C int) (int, []int) {
	n := len(V)
	A := make([][]int, n+1) // solution placeholder
	for i := range A {
		A[i] = make([]int, C+1)
	}
	for i := 1; i < n+1; i++ {
		for c := 0; c < C+1; c++ {
			if S[i-1] > c {
				A[i][c] = A[i-1][c]
			} else {
				A[i][c] = max(A[i-1][c], A[i-1][c-S[i-1]]+V[i-1])
			}
		}
	}
	// Reconstruct objects stuffed
	J := make([]int, 0)
	c := C
	for i := n; i > 0; i-- {
		if S[i-1] <= c && A[i-1][c-S[i-1]]+V[i-1] >= A[i-1][c] {
			J = append(J, i)
			c -= S[i-1]
		}
	}
	return A[n][C], J
}

func main() {
	var n, C int
	// Parse input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	str := make([]string, 0)
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	n = pop(&str)                          // number of objects
	C = pop(&str)                          // capacity
	V, S := make([]int, n), make([]int, n) // values and sizes
	for i := range V {
		V[i] = pop(&str)
	}
	for i := range S {
		S[i] = pop(&str)
	}
	w, J := stuffem(V, S, C)
	fmt.Printf("Stuff objects %v into the sack for %d\n", J, w)
}
