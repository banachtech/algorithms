package main

import (
	"fmt"
)

func min(a ...int) int {
	tmp := a[0]
	for _, v := range a {
		if v < tmp {
			tmp = v
		}
	}
	return tmp
}

func mindist(s1, s2 string) int {
	m := len(s1)
	n := len(s2)
	// sub problem solution placeholder
	P := make([][]int, m+1)
	// initialise base case
	// 0th row and 0th column represent empty first and second string
	for i := range P {
		P[i] = make([]int, n+1)
		P[i][0] = i
	}
	for i := range P[0] {
		P[0][i] = i
	}
	// cost function
	// if ith char in s1 and jth char in s2 differ set cost = 1
	a := make([][]int, m)
	for i := range a {
		a[i] = make([]int, n)
		for j := range a[i] {
			if s1[i] != s2[j] {
				a[i][j] = 1
			}
		}
	}
	// compute P recursively
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			// pick min of 3 cases: chars differ (2 ops required) or one of them is a space
			P[i][j] = min(P[i-1][j-1]+a[i-1][j-1], P[i-1][j]+1, P[i][j-1]+1)
		}
	}
	return P[m][n]
}

func main() {
	var s1, s2 string
	fmt.Scanln(&s1)
	fmt.Scanln(&s2)
	d := mindist(s1, s2)
	fmt.Println(d)
}
