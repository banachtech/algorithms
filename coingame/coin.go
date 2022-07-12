package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func max(x ...int) int {
	out := x[0]
	for i := range x {
		if x[i] > out {
			out = x[i]
		}
	}
	return out
}

func diff(x []int, i int) []int {
	x[0], x[i] = x[i], x[0]
	return x[1:]
}

func play(c []int, i int) int {
	n := len(c)
	if i == 0 && n == 2 {
		return max(c...)
	}
	if i == 1 && n == 1 {
		return c[0]
	}
	k := (i + 1) % 2
	return max(c[0]+play(diff(c[1:], play(c[1:], k)), i), c[n-1]+play(diff(c[:n-1], play(c[:n-1], k)), i))
}

func next1(x []int) ([]int, int) {
	x = diff(x, play2(x))
	return x, play1(x)
}

func next2(x []int) ([]int, int) {
	x = diff(x, play2(x))
	return x, play1(x)
}

func play1(c []int) int {
	n := len(c)
	// Terminal condition
	if n == 2 {
		if c[1] > c[0] {
			payout += c[1]
			return 1
		} else {
			payout += c[0]
			return 0
		}
	}
	x0, i0 := next1(diff(c, 0))
	x1, i1 := next1(diff(c, n-1))
	if c[n-1]+x1[i1] > c[0]+x0[i0] {
		payout += c[n-1]
	} else {
		payout += c[0]
	}
}

func play2(c []int) int {
	n := len(c)
	// Terminal condition
	if n == 1 {
		return 0
	}
	x0, i0 := next2(diff(c, 0))
	x1, i1 := next2(diff(c, n-1))
	if c[n-1]+x1[i1] > c[0]+x0[i0] {
		return n - 1
	} else {
		return 0
	}
}

var payout int

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	c := make([]int, n)
	for i := range c {
		c[i] = rand.Intn(10) * 2
	}
	fmt.Printf("coins are: %v\n", c)
	play1(c)
	fmt.Println(payout)
}
