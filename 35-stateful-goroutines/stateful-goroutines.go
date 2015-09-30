package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func ds(reads <-chan *readOp, writes <-chan *writeOp, ops *uint64) {
	var state = make(map[int]int)
	for {
		select {
		case read := <-reads:
			read.resp <- state[read.key]
		case write := <-writes:
			state[write.key] = write.val
			write.resp <- true
		}
		atomic.AddUint64(ops, 1)
	}
}

func reader(reads chan<- *readOp) {
	for {
		read := &readOp{
			key:  rand.Intn(5),
			resp: make(chan int)}
		reads <- read
		<-read.resp
	}
}

func writer(writes chan<- *writeOp) {
	for {
		write := &writeOp{
			key:  rand.Intn(5),
			val:  rand.Intn(100),
			resp: make(chan bool)}
		writes <- write
		<-write.resp
	}
}

func main() {
	var ops uint64 = 0
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	go ds(reads, writes, &ops)

	for r := 0; r < 100; r++ {
		go reader(reads)
	}

	for w := 0; w < 10; w++ {
		go writer(writes)
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
