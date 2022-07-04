// Tower of Hanoi
package main

import (
	"fmt"
	"os"
	"strconv"
)

func hanoi(n, i, j int) {
	if n == 1 {
		fmt.Printf("%c\t  %c\n", tower[i], tower[j])
		moves++
		return
	}
	hanoi(n-1, i, 3-i-j)
	fmt.Printf("%c\t  %c\n", tower[i], tower[j])
	moves++
	hanoi(n-1, 3-i-j, j)
}

var moves int
var tower map[int]rune

func main() {
	var n int
	tower = make(map[int]rune)
	tower[0] = 'A'
	tower[1] = 'B'
	tower[2] = 'C'
	if len(os.Args) < 2 {
		n = 2
	} else {
		n, _ = strconv.Atoi(os.Args[1])
	}
	fmt.Printf("From\t To\n")
	fmt.Printf("-----------\n")
	hanoi(n, 0, 2)
	fmt.Printf("-----------\n")
	fmt.Printf("Took %d moves!\n", moves)
	fmt.Printf("-----------\n")
}
