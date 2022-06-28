package main

import (
	"fmt"
	"time"
)

func gcd1(a, b int) int {
	for {
		if a <= 0 || b <= 0 {
			break
		}
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return max(a, b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func gcd(a, b int) int {
	if a == 0 || b == 0 {
		return max(a, b)
	}
	return gcd(b, a%b)
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	start := time.Now()
	fmt.Println(lcm(n, m))
	fmt.Println(time.Since(start))
}
