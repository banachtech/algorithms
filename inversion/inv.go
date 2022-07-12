package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func cross(a, b []int) int {
	n := 0
	for _, v := range a {
		for _, w := range b {
			if v > w {
				n++
			}
		}
	}
	return n
}

func inv(x []int) int {
	n := len(x)
	if n <= 1 {
		return 0
	}
	m := n / 2
	left := x[:m]
	right := x[m:]
	return inv(left) + inv(right) + cross(left, right)
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
	y := make([]int, n)
	for i := range y {
		y[i] = pop(&str)
	}

	fmt.Printf("%d\n", inv(y))
}
