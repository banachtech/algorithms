package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func numStops(location, distance, tank int, stops []int) int {
	if location+tank >= distance {
		return 0
	}
	if location+tank < stops[0] {
		return math.MinInt
	}
	lastStop := location
	for len(stops) > 0 && stops[0]-location <= tank {
		lastStop = stops[0]
		stops = stops[1:]
	}
	return 1 + numStops(lastStop, distance, tank, stops)
}

func main() {
	// Scan input
	var d, m, n int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	str := make([]string, 0)
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	d, _ = strconv.Atoi(str[0])
	m, _ = strconv.Atoi(str[1])
	n, _ = strconv.Atoi(str[2])
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i], _ = strconv.Atoi(str[i+3])
	}
	stops := numStops(0, d, m, s)
	if stops < 0 {
		stops = -1
	}
	fmt.Println(stops)
}
