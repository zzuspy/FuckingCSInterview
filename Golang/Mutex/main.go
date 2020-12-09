package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

const (
	mutexLockedd = iota
)

func main() {
	for i := 1; i <= 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	x := sync.RWMutex{}
	time.Sleep(time.Second)
}
