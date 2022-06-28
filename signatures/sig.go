package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func minslice(x []int) int {
	out := x[0]
	for _, v := range x {
		if v < out {
			out = v
		}
	}
	return out
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	str := make([]string, 0)
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	n, _ := strconv.Atoi(str[0])
	st := make([]int, n)
	ed := make([]int, n)
	for i := range st {
		st[i], _ = strconv.Atoi(str[2*i+1])
		ed[i], _ = strconv.Atoi(str[2*i+2])
	}
	var pts []int
	m := 0
	x := ed
	for len(x) > 0 {
		m = minslice(x)
		pts = append(pts, m)
		x = nil
		for i := range st {
			if st[i] > m {
				x = append(x, ed[i])
			}
		}
	}
	fmt.Println(len(pts))
	for _, v := range pts {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")
}
