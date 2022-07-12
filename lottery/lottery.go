package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func pop(z *[]string) int {
	tmp, _ := strconv.Atoi((*z)[0])
	*z = (*z)[1:]
	return tmp
}

func main() {
	var n, m int
	// Parse input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	str := make([]string, 0)
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	n = pop(&str)
	m = pop(&str)

	// collect segments
	// store the left edge coordinate along with +1 and
	// right edge coordinate along with -1 in a 2d slice
	seg := make([][]int, 2*n)
	for i := range seg {
		seg[i] = make([]int, 2)
		x := pop(&str)
		if i%2 == 0 {
			seg[i][0], seg[i][1] = x, 1
		} else {
			seg[i][0], seg[i][1] = x+1, -1
		}
	}
	// sort in descending order along first dimension
	sort.SliceStable(seg, func(i, j int) bool { return seg[i][0] >= seg[j][0] })

	// collect points
	p := make([]int, m)
	for i := range p {
		p[i] = pop(&str)
	}
	// make a sorted copy
	q := make([]int, m)
	copy(q, p)
	sort.Ints(q)
	// store number of coverings
	counter := make(map[int]int)
	// flag to avoid work on repeats
	seen := make(map[int]bool)

	// key insight is that for point p2 to the right of
	// point p1, the count is a cumulative adjustment to count
	// for point p1; this avoids extra looping
	count := 0
	for _, v := range q {
		if !seen[v] {
			// as long as the left or right of segment is
			// to the left of the point
			for len(seg) > 0 && seg[len(seg)-1][0] <= v {
				// -1 takes care of points that fall beyond
				// the right edge of a segment
				count += seg[len(seg)-1][1]
				// remove the segment
				seg = seg[:len(seg)-1]
			}
			counter[v] = count
		}
	}
	// print counts
	for _, v := range p {
		fmt.Printf("%d ", counter[v])
	}
	fmt.Printf("\n")

}
