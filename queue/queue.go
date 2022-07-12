package main

import (
	"fmt"
	"math/rand"
)

type queue []interface{}

func NewQ() queue {
	return make([]interface{}, 0)
}

func (q *queue) enqueue(x interface{}) {
	*q = append(*q, x)
}

func (q *queue) dequeue() interface{} {
	if len(*q) == 0 {
		return nil
	}
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func main() {
	x := make([]int64, 10)
	q := NewQ()

	fmt.Println("enqueuing ...")
	for i := range x {
		x[i] = int64(rand.Int())
		fmt.Println(x[i])
		q.enqueue(x[i])
	}
	fmt.Println("dequeuing ...")
	for {
		y := q.dequeue()
		if y != nil {
			fmt.Println(y)
		} else {
			break
		}
	}
}
