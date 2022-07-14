package main

import "fmt"

func min(a ...int) int {
	tmp := a[0]
	for _, v := range a {
		if v < tmp {
			tmp = v
		}
	}
	return tmp
}

func main() {
	var n, m int
	fmt.Scanf("%d", &n)
	c := []int{1, 2, 1, 1}
	if n < 5 {
		m = c[n-1]
	} else {
		x := make([]int, n)
		copy(x, c)
		for i := 4; i < n; i++ {
			x[i] = min(x[i-4], x[i-3], x[i-1]) + 1
		}
		m = x[n-1]
	}
	fmt.Println(m)
}
