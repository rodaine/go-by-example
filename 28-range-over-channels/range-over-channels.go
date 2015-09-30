package main

import "fmt"

func main() {
	q := make(chan int, 2)
	q <- 1
	q <- 2
	close(q)

	for elem := range q {
		fmt.Println(elem)
	}
}
