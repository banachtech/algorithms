package main

import (
	"fmt"
)

func length(x int) int {
	n := 1
	for x/10 > 0 {
		n++
		x = x / 10
	}
	return n
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func pow10(n int) int {
	if n == 0 {
		return 1
	}
	out := 1
	for i := 1; i <= n; i++ {
		out = out * 10
	}
	return out
}

func karatsuba(x, y int) int {
	n := max(length(x), length(y))
	if n < 2 {
		return x * y
	}
	m := n/2 + n%2
	l := pow10(m)
	a := x / l
	b := x % l
	c := y / l
	d := y % l
	p := karatsuba(a, c)
	q := karatsuba(a+b, c+d)
	r := karatsuba(b, d)
	return pow10(m*2)*p + l*(q-p-r) + r
}

func main() {
	x := 123
	y := 456
	fmt.Printf("Want %d, got %d\n", x*y, karatsuba(x, y))
}
