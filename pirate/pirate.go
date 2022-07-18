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

// compute max stuffed value and stuffed objects
func stuffem(S []int, C int) (int, []int) {
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
	// stuffed weights
	J := make([]int, 0)
	c := C
	for i := n; i >= 1; i-- {
		if S[i-1] <= c && A[i-1][c-S[i-1]]+S[i-1] > A[i-1][c] {
			J = append(J, i-1)
			c = c - S[i-1]
		}
	}
	return A[n][C], J
}

// remove indices y from x
func remove(x, y []int) []int {
	for _, j := range y {
		x[0], x[j] = x[j], x[0]
		x = x[1:]
	}
	return x
}

func main() {
	var n int
	var can int
	// Parse input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	str := make([]string, 0)
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	n = pop(&str)
	fmt.Println(n)          // number of integers
	items := make([]int, n) // integers
	for i := range items {
		items[i] = pop(&str)
		fmt.Println(items)
	}
	totalValue := sum(items) // total loot
	// prelim checks
	if totalValue%3 == 0 && len(items) >= 3 {
		fmt.Println("inside if")
		can = 1
		share := totalValue / 3
		for i := 0; i < 3; i++ {
			fmt.Printf("inside iteration %d\n", i)
			val, stuff := stuffem(items, share)
			fmt.Printf("%d %v\n", val, stuff)
			if val < share {
				can = 0
				break
			}
			items = remove(items, stuff)
		}
	}
	fmt.Println(can)
}
