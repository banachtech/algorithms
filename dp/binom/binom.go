package main

import (
	"fmt"
	"time"
)

func binom(n, k int) int {
	b := make([][]int, n+1)
	for i := range b {
		b[i] = make([]int, n+1)
		b[i][0] = 1
		b[i][i] = 1
	}
	for i := 2; i < n+1; i++ {
		for j := 1; j < i; j++ {
			b[i][j] = b[i-1][j-1] + b[i-1][j]
		}
	}
	return b[n][k]
}

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	st := time.Now()
	fmt.Printf("%dC%d = %d in %v\n", n, k, binom(n, k), time.Since(st))
}
