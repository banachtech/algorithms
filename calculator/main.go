package main

import "fmt"

func min(a ...int) int {
	tmp := a[0]
	for _, v := range a {
		if tmp > v {
			tmp = v
		}
	}
	return tmp
}

func calculator(n int) []int {
	table := make([]int, n+1)
	for i := 2; i < n+1; i++ {
		table[i] = 1 + table[i-1]
		if i%2 == 0 {
			table[i] = min(table[i], 1+table[i/2])
		}
		if i%3 == 0 {
			table[i] = min(table[i], 1+table[i/3])
		}
	}
	var ops []int
	for n > 1 {
		ops = append(ops, n)
		if table[n] == 1+table[n-1] {
			n = n - 1
		} else if n%2 == 0 && table[n] == 1+table[n/2] {
			n = n / 2
		} else if n%3 == 0 && table[n] == 1+table[n/3] {
			n = n / 3
		}
	}
	ops = append(ops, 1)
	return ops
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	o := calculator(n)
	fmt.Println(len(o) - 1)
	for i := len(o) - 1; i >= 0; i-- {
		fmt.Printf("%d ", o[i])
	}
	fmt.Printf("\n")
}
