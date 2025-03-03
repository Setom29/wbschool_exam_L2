package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) *
				time.Millisecond)
		}
		close(c)
	}()
	return c
}
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}

	}()
	return c
}

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}

// package main

// import "fmt"

// func main() {
// 	s := make([]int, 0)
// 	var oldCap int
// 	for i := 0; i < 100000; i++ {
// 		s = append(s, i)
// 		if cap(s) != oldCap {
// 			if oldCap != 0 {
// 				fmt.Println(i, oldCap, float64(cap(s))/float64(oldCap))
// 			}
// 			oldCap = cap(s)
// 		}
// 	}

// }
