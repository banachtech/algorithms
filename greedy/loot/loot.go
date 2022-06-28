package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Slice struct {
	w   []float64
	c   []float64
	ids []int
}

func NewSlice(a, b []float64) Slice {
	s := Slice{a, b, make([]int, len(a))}
	for i := range s.ids {
		s.ids[i] = i
	}
	return s
}

func (s Slice) Len() int {
	return len(s.w)
}

func (s Slice) Swap(i, j int) {
	s.c[i], s.c[j] = s.c[j], s.c[i]
	s.w[i], s.w[j] = s.w[j], s.w[i]
}

func (s Slice) Less(i, j int) bool {
	if s.c[j] < s.c[i] {
		return true
	}
	return false
}

// Find allocations fi that maximizes sum(ci*fi)
// subject to sum(wi*fi) <= W
// Assumption: c is sorted in increasing order
func alloc(w, c []float64, W float64) float64 {
	var x, v float64
	slack := W
	for i := range c {
		x = math.Min(w[i], slack)
		slack -= x
		v += x * c[i]
		if slack <= 0 {
			break
		}
	}
	return v
}

func max(m []float64) (int, float64) {
	out := 0.0
	j := 0
	for i, v := range m {
		if v > out {
			out = v
			j = i
		}
	}
	return j, out
}

func alloc1(w, c []float64, W float64) float64 {
	x := 0.0
	val := 0.0
	for W > 0 {
		i, v := max(c)
		x = math.Min(w[i], W)
		W -= x
		val += x * v
		c[0], c[i] = c[i], c[0]
		w[0], w[i] = w[i], w[0]
		c = c[1:]
		w = w[1:]
	}
	return val
}

func main() {
	// Scan input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	s := make([]string, 0)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}
	// Total weight
	W, _ := strconv.ParseFloat(s[1], 64)
	n, _ := strconv.Atoi(s[0])
	// Parse weights and cost
	w := make([]float64, n)
	c := make([]float64, n)
	s = s[2:]
	for j := 0; j < n; j++ {
		c[j], _ = strconv.ParseFloat(s[2*j], 64)
		w[j], _ = strconv.ParseFloat(s[2*j+1], 64)
		c[j] /= w[j]
	}
	fmt.Println(c)
	fmt.Println(w)
	// Compute total value
	val := alloc1(w, c, W)
	fmt.Printf("value: %.4f\n", val)
}
