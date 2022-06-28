package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	var r int
	for s.Scan() {
		r, _ = strconv.Atoi(s.Text())
	}
	n := getprizes(r, 1)
	for _, v := range n {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")
}

func getprizes(n, minval int) []int {
	if n < 2*minval+1 {
		return []int{n}
	}
	return append(getprizes(n-minval, minval+1), minval)

}
