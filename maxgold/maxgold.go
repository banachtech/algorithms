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

// compute max stuffed value
func stuffem(S []int, C int) int {
	n := len(S)
	A := make([][]int, n+1) // solution placeholder
	for i := range A {
		A[i] = make([]int, C+1)
	}
	for i := 1; i < n+1; i++ {
		for c := 0; c < C+1; c++ {
			if S[i-1] > c {
				A[i][c] = A[i-1][c]
			} else {
				A[i][c] = max(A[i-1][c], A[i-1][c-S[i-1]]+S[i-1])
			}
		}
	}
	return A[n][C]
}

func main() {
	var n, W int
	// Parse input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	str := make([]string, 0)
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	W = pop(&str)       // capacity
	n = pop(&str)       // number of objects
	w := make([]int, n) // weights of the bars
	for i := range w {
		w[i] = pop(&str)
	}
	fmt.Println(stuffem(w, W))
}
