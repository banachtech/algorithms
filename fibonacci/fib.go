package main

import (
	"fmt"
	"log"
	"math/big"
	"time"
)

// Standard version
func fibStandard(n int) int {
	if n < 2 {
		return n
	}
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

// Standard fibonacci with arbitrary precision
func bigfib(m int64) *big.Int {
	var i int64
	if m < 0 {
		return big.NewInt(m)
	}
	a, b := big.NewInt(0), big.NewInt(1)
	for i = 0; i < m; i++ {
		a, b = b, a.Add(a, b)
	}
	return a
}

// Closure version with arbitrary precision
func bigfibClosure(m int64) *big.Int {
	var i int64
	n := big.NewInt(0)
	if m < 2 {
		n.SetInt64(m)
		return n
	}
	f := func() func() *big.Int {
		x1, x2 := big.NewInt(0), big.NewInt(1)
		return func() *big.Int {
			x := big.NewInt(0)
			x.Add(x1, x2)
			x1, x2 = x2, x
			return x
		}
	}()
	for i = 0; i < m-1; i++ {
		n = f()
	}
	return n
}

// Fibonacci with memoization technique
func fibMemoized(n int) int {
	tbl := make(map[int]int)
	return memoize(n, tbl)
}
func memoize(m int, tbl map[int]int) int {
	if m < 2 {
		return m
	}
	if v, ok := tbl[m]; ok {
		return v
	}
	tbl[m] = memoize(m-2, tbl) + memoize(m-1, tbl)
	return tbl[m]
}

// Fibonacchi modulo m
func fibmod(n, m int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, (a+b)%m
	}
	return a
}

// Last digit of Fibonacchi
func lastfib(m int) int {
	if m < 0 {
		return m
	}
	a, b := 0, 1
	for i := 0; i < m; i++ {
		a = (a + b) % 10
		a, b = b, a
	}
	return a
}

// Pisano period
func fibperiod(m int) int {
	a, b := 0, 1
	p := 0
	for {
		a, b = b, (a+b)%m
		p++
		if a == 0 && b == 1 {
			return p
		}
	}
}

// Last digit of sum of fibonacci numbers
// Use the fact that S(n) = F(n+2)-1
func lastdigitsum(n int) int {
	a, b := 0, 1
	for i := 0; i < n+2; i++ {
		a, b = b, (a+b)%10
	}
	if a == 0 {
		return 9
	}
	return (a - 1) % 10
}

// Fast exponentiation using squaring
func exp(x, n int) int {
	if n == 0 {
		return 1
	}
	if n%2 == 1 {
		return x * exp(x*x, (n-1)/2)
	} else {
		return exp(x*x, n/2)
	}
}

// Fast matrix exponentiation
func expMat(x [][]int, n int) [][]int {
	if n == 0 {
		y := make([][]int, len(x))
		for i := range y {
			y[i] = make([]int, len(x[0]))
			y[i][i] = 1
		}
		return y
	}
	xx := matmul(x, x)
	if n%2 == 1 {
		return matmul(x, expMat(xx, (n-1)/2))
	} else {
		return expMat(xx, n/2)
	}
}

// Matrix multiplication
func matmul(x, y [][]int) [][]int {
	if len(x[0]) != len(y) {
		log.Fatal("dimension mismatch")
		return nil
	}
	z := make([][]int, len(x))
	for i := range x {
		z[i] = make([]int, len(y[0]))
		for j := 0; j < len(y[0]); j++ {
			z[i][j] = 0
			for k := range x[i] {
				z[i][j] += x[i][k] * y[k][j]
			}
		}
	}
	return z
}

// Matrix exponentiation mod m
func modexpMat(x [][]int, n, m int) [][]int {
	if n == 0 {
		y := make([][]int, len(x))
		for i := range y {
			y[i] = make([]int, len(x[0]))
			y[i][i] = 1
		}
		return y
	}
	xx := modmatmul(x, x, m)
	if n%2 == 1 {
		return modmatmul(x, modexpMat(xx, (n-1)/2, m), m)
	} else {
		return modexpMat(xx, n/2, m)
	}
}

// Matrix multiplication mod m
func modmatmul(x, y [][]int, m int) [][]int {
	z := make([][]int, len(x))
	for i := range x {
		z[i] = make([]int, len(y[0]))
		for j := 0; j < len(y[0]); j++ {
			z[i][j] = 0
			for k := range x[i] {
				z[i][j] += x[i][k] * y[k][j]
			}
			z[i][j] = z[i][j] % m
		}
	}
	return z
}

/*
func main() {
	var n int
	fmt.Scanf("%d %d", &n)
	// Compute last digit of sum of fibonacci series of length n fast
	// Find the period of Fn mod 10
	p := fibperiod(10)
	// Use the periodicity to reduce the n+2 in F(n+2) to manageable size
	d := (fibStandard((n+2)%p) - 1) % 10
	start := time.Now()
	fmt.Printf("Last digit of S_%d is %d\nTook %v seconds\n", n, d, time.Since(start))

	// Compute last digit of sum of fibonacci series of length n
	// using matrix exponentiation from the relationship
	// |F(n-1)	F(n)| = |0 1|^n
	// |F(n)  F(n+1)| = |1 1|
	x := make([][]int, 2)
	for i := range x {
		x[i] = make([]int, 2)
	}
	x[0][0], x[0][1] = 0, 1
	x[1][0], x[1][1] = 1, 1
	start = time.Now()
	y := modexpMat(x, n+2, 10)
	fmt.Printf("Exponentiation took %v\n", time.Since(start))
	fmt.Printf("%d %d\n%d %d\n", y[0][0], y[0][1], y[1][0], y[1][1])
	// (0,1) and (1,0) entries of y are F(n)
	fmt.Printf("Last digit of S_%d is %d\n", n, (y[0][1]-1)%10)
}
*/

func main() {
	var m, n int
	fmt.Scanf("%d %d", &m, &n)
	// Compute last digit of partial sum F(m)+F(m+1)+...+F(n)
	// Note: 60 is the pisano period for mod 10
	start := time.Now()
	d := (fibStandard((n+2)%60)%10 - fibStandard((m+1)%60)%10) % 10
	if d < 0 {
		d += 10
	}
	fmt.Printf("last digit %d, took %v\n", d, time.Since(start))
}
