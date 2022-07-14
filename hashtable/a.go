package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func sum2(x []int, y int) int {
	yes := make(map[int]bool)
	count := 0
	for i := range x {
		if !yes[x[i]] {
			yes[x[i]] = true
		}
	}
	for i := range x {
		if yes[y-x[i]] {
			count++
		}
	}
	return count
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	m, _ := strconv.Atoi(os.Args[2])
	s, _ := strconv.Atoi(os.Args[3])
	x := make([]int, s)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range x {
		x[i] = r.Intn(n)
	}
	p := float64(sum2(x, m)) / float64(s)
	fmt.Println(p)
}
