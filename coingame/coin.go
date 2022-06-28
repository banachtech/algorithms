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

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	c := make([]int, n)
	for i := range c {
		c[i] = rand.Intn(10) * 2
	}
	fmt.Printf("coins are: %v\n", c)
	fmt.Println(play(c, 1))
}
