package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

const iterationsNum = 7
const goroutinesNum = 5

func startWorker(in int, wg *sync.WaitGroup) {
	defer wg.Done() //уменьшает счетчик на 1
	for j := 0; j < iterationsNum; j++ {
		fmt.Println(formatWork(in, j))
		runtime.Gosched()
	}
}

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < goroutinesNum; i++ {
		wg.Add(1)
		go startWorker(i, wg)
	}
	time.Sleep(time.Millisecond)
	wg.Wait()

}

func formatWork(in, j int) string {
	return fmt.Sprintln(strings.Repeat("  ", in), "█",
		strings.Repeat("  ", goroutinesNum-in),
		"th", in,
		"iter", j, strings.Repeat("■", j))
}
