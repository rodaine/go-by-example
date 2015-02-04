package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func main() {
	a, b := 1, 2
	res := plus(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, res)
}
