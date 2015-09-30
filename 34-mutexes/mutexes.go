package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func readWorker(m *sync.Mutex, s map[int]int, o *uint64) {
	total := 0
	for {
		key := rand.Intn(5)
		m.Lock()
		total += s[key]
		m.Unlock()
		atomic.AddUint64(o, 1)
		runtime.Gosched()
	}
}

func writeWorker(m *sync.Mutex, s map[int]int, o *uint64) {
	for {
		key := rand.Intn(5)
		val := rand.Intn(100)
		m.Lock()
		s[key] = val
		m.Unlock()
		atomic.AddUint64(o, 1)
		runtime.Gosched()
	}
}

func main() {
	state := make(map[int]int)
	mutex := sync.Mutex{}
	var ops uint64 = 0

	for r := 0; r < 100; r++ {
		go readWorker(&mutex, state, &ops)
	}

	for w := 0; w < 10; w++ {
		go writeWorker(&mutex, state, &ops)
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
