package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// check if x is in the sorted slice y
// recursive implementation
func recursiveIn(x int, y []int) int {
	if len(y) == 0 {
		return -1
	}
	m := len(y) / 2
	if m > 0 {
		if x == y[m] {
			return m + offset
		}
		if x < y[m] {
			return recursiveIn(x, y[:m])
		}
		if x > y[m] {
			offset += m
			return recursiveIn(x, y[m:])
		}
	}
	if x == y[m] {
		return m + offset
	} else {
		return -1
	}
}

func searchSortedFirst(x int, y []int) int {
	curId := -1
	nextId := -1
	for {
		offset = 0
		nextId = recursiveIn(x, y)
		if nextId != -1 {
			curId = nextId
			y = y[:nextId]
		} else {
			break
		}
	}
	return curId
}

var offset int

func pop(z *[]string) int {
	tmp, _ := strconv.Atoi((*z)[0])
	*z = (*z)[1:]
	return tmp
}

func isMajority(y []int) int {
	var n, m, q int
	sort.Ints(y)
	n = len(y)
	if n == 1 {
		return 1
	}
	q = n / 2
	s := make(map[int]bool)
	for _, v := range y {
		if !s[v] {
			m = count(v, y)
			fmt.Printf("%d occurs %d times\n", v, m)
			if m > q {
				return 1
			}
			s[v] = true
		}
	}
	return 0
}

func count(x int, y []int) int {
	i := searchSortedFirst(x, y)
	if i == -1 || len(y) == 0 {
		return 0
	}
	return 1 + count(x, y[i+1:])
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
	n = pop(&str)
	y := make([]int, n)
	for i := range y {
		y[i] = pop(&str)
	}
	fmt.Printf("%d\n", isMajority(y))
}

func in(x int, y []int) bool {
	if len(y) == 0 {
		return false
	}
	m := len(y) / 2
	if m > 0 {
		if x == y[m] {
			return true
		}
		if x < y[m] {
			return in(x, y[:m])
		}
		if x > y[m] {
			return in(x, y[m:])
		}
	}
	if x == y[m] {
		return true
	} else {
		return false
	}
}
