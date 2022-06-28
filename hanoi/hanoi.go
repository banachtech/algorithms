// Tower of Hanoi
package main

import "fmt"

func hanoi(n, i, j int) {
	if n == 1 {
		fmt.Printf("move disc from %d to %d\n", i, j)
		moves++
		return
	}
	hanoi(n-1, i, 3-i-j)
	fmt.Printf("move disc from %d to %d\n", i, j)
	moves++
	hanoi(n-1, 3-i-j, j)
}

var moves int

func main() {
	hanoi(6, 0, 2)
	fmt.Println(moves)
}
