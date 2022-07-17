package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func pop(z *[]string) int {
	tmp, _ := strconv.Atoi((*z)[0])
	*z = (*z)[1:]
	return tmp
}

func max(a ...int) int {
	tmp := a[0]
	for _, v := range a {
		if v > tmp {
			tmp = v
		}
	}
	return tmp
}

func dpsq3(a, b, c []int) int {
	n, m, l := len(a), len(b), len(c)
	X := make([][][]int, n+1)
	for i := 0; i < n+1; i++ {
		X[i] = make([][]int, m+1)
		for j := 0; j < m+1; j++ {
			X[i][j] = make([]int, l+1)
			for k := 0; k < l+1; k++ {
				if i == 0 || j == 0 || k == 0 {
					X[i][j][k] = 0
				} else {
					if (a[i-1] == b[j-1]) && (a[i-1] == c[k-1]) && (b[j-1] == c[k-1]) {
						X[i][j][k] = 1 + X[i-1][j-1][k-1]
					} else {
						X[i][j][k] = max(X[i-1][j][k], X[i][j-1][k], X[i][j][k-1])
						X[i][j][k] = max(X[i][j][k], X[i-1][j-1][k], X[i][j-1][k-1], X[i-1][j][k-1])
					}
				}
			}

		}
	}
	return X[n][m][l]
}

func main() {
	var n, m, l int
	// Parse input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	str := make([]string, 0)
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	n = pop(&str)
	a := make([]int, n)
	for i := range a {
		a[i] = pop(&str)
	}
	m = pop(&str)
	b := make([]int, m)
	for i := range b {
		b[i] = pop(&str)
	}
	l = pop(&str)
	c := make([]int, l)
	for i := range c {
		c[i] = pop(&str)
	}
	fmt.Println(dpsq3(a, b, c))
}

func dpsq(a, b []int) int {
	n, m := len(a), len(b)
	X := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		X[i] = make([]int, m+1)
		for j := 0; j < m+1; j++ {
			if i == 0 || j == 0 {
				X[i][j] = 0
			} else {
				if a[i-1] == b[j-1] {
					X[i][j] = 1 + X[i-1][j-1]
				} else {
					X[i][j] = max(X[i-1][j], X[i][j-1])
				}
			}
		}
	}
	return X[n][m]
}

func naive(a, b []int) int {
	n, m := len(a), len(b)
	if n == 0 || m == 0 {
		return 0
	}
	if (n == 1 && in(a[0], b)) || (m == 1 && in(b[0], a)) {
		return 1
	}
	if a[n-1] == b[m-1] {
		return naive(a[:n-1], b[:m-1]) + 1
	} else {
		return max(naive(a[:n-1], b), naive(a, b[:m-1]))
	}
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
