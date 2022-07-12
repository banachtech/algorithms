package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

func dist(x, y []int) float64 {
	return math.Hypot(float64(x[0]-y[0]), float64(x[1]-y[1]))
}

func minf(x []float64) float64 {
	d := math.MaxFloat64
	for i := range x {
		d = math.Min(d, x[i])
	}
	return d
}

func mindist7(x [][]int, d float64) float64 {
	tmp := d
	n := len(x)
	// sort by y coordinates
	sort.SliceStable(x, func(i, j int) bool { return x[i][1] <= x[j][1] })
	for i := 0; i < n; i++ {
		for j := i + 1; j < n && math.Abs(float64(x[j][1]-x[i][1])) < tmp; j++ {
			tmp = math.Min(tmp, dist(x[i], x[j]))
		}
	}
	return tmp
}

func mindist(x [][]int) float64 {
	n := len(x)
	// brute force if n <= 3
	switch n {
	case 0, 1:
		return math.MaxFloat64
	case 2:
		return dist(x[0], x[1])
	case 3:
		d1, d2, d3 := dist(x[0], x[1]), dist(x[0], x[2]), dist(x[1], x[2])
		return minf([]float64{d1, d2, d3})
	}
	// mid point
	m := n / 2
	mx := x[m][0]
	// min of min distances for regions left and right of mid point
	d := math.Min(mindist(x[:m]), mindist(x[m+1:]))

	// points around divider with x coordinates within distance d
	z := make([][]int, 0)
	for i := range x {
		if math.Abs(float64(x[i][0]-mx)) < d {
			z = append(z, x[i])
		}
	}
	d = math.Min(d, mindist7(z, d))
	return d
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
	n = pop(&str)
	// collect points
	p := make([][]int, n)
	for i := range p {
		p[i] = make([]int, 2)
		p[i][0], p[i][1] = pop(&str), pop(&str)
	}
	// sort points by x co-ordinates
	sort.SliceStable(p, func(i, j int) bool { return p[i][0] <= p[j][0] })

	// compute min distance
	d := mindist(p)

	fmt.Printf("%f\n", d)
}

func fakedata(n int) [][]int {
	p := make([][]int, n)
	for i := range p {
		p[i] = make([]int, 2)
		p[i][0] = rand.Intn(100000)
		p[i][1] = rand.Intn(100000)
	}
	return p
}
