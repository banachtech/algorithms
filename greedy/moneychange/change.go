package main

import "fmt"

func main() {
	var n, d int
	fmt.Scanf("%d", &n)
	a := []int{1, 5, 10}
	count := 0
	for {
		if n == 0 {
			break
		}
		d = 0
		for _, v := range a {
			if v <= n && v > d {
				d = v
			}
		}
		n = n - d
		count++
	}
	fmt.Println(count)
}
