package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i]+p[j] <= p[j]+p[i] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	// Parse input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	str := make([]string, 0)
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	str = str[1:]
	sort.Sort(sort.Reverse(StringSlice(str)))
	fmt.Println(strings.Join(str, ""))
}
