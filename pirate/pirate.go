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

func sum(x []int) int {
	tmp := 0
	for _, v := range x {
		tmp += v
	}
	return tmp
}

func pop(z *[]string) int {
	tmp, _ := strconv.Atoi((*z)[0])
	*z = (*z)[1:]
	return tmp
}

func main() {
	var n int
	// Parse input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	str := make([]string, 0)
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	n = pop(&str)       // number of integers
	x := make([]int, n) // integers
	for i := range x {
		x[i] = pop(&str)
	}
	can := split(x)
	if can {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

func split(x []int) bool {
	// prelim checks
	y := sum(x)
	n := len(x)
	if len(x) < 3 {
		return false
	}
	if y%3 != 0 {
		return false
	}
	seen := make(map[int]bool)
	seen[n-1] = true
	s := make([]int, 3)
	s[0] = x[n-1]
	return foo(x, s, seen, 0, n-1)
}

func foo(a, s []int, seen map[int]bool, is, ia int) bool {
	share := sum(a) / 3
	if s[is] == share {
		if is == 1 {
			return true
		}
		return foo(a, s, seen, is+1, ia)
	}
	for i := ia; i >= 0; i-- {
		if !seen[i] {
			if s[is]+a[i] <= share {
				seen[i] = true
				s[is] += a[i]
				tmp := foo(a, s, seen, is, i-1)
				if tmp {
					return true
				}
				seen[i] = false // backtrack
				s[is] -= a[i]
			}
		}
	}
	return false
}
