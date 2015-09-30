package main

import "time"
import "fmt"

func sleeper(c chan<- string, id string) {
	time.Sleep(time.Second * 2)
	c <- fmt.Sprint("result ", id)
}

func timeout(c <-chan string, id string, to time.Duration) {
	select {
	case res := <-c:
		fmt.Println(res)
	case <-time.After(to):
		fmt.Println("timeout", id)
	}
}

func main() {
	c1 := make(chan string, 1)
	go sleeper(c1, "1")
	timeout(c1, "1", time.Second*1)

	c2 := make(chan string, 1)
	go sleeper(c2, "2")
	timeout(c2, "2", time.Second*3)
}
