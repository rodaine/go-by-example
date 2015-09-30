package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func counter(ops *uint64) {
	for {
		atomic.AddUint64(ops, 1)
		runtime.Gosched()
	}
}

func main() {
	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go counter(&ops)
	}

	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
