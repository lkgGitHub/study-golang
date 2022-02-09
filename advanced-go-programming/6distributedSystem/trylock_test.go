package main

import (
	"sync"
	"testing"
)

/*
尝试加锁，加锁成功执行后续流程，如果加锁失败的话也不会阻塞，而会直接返回加锁的结果。

在单机系统中，trylock并不是一个好选择。因为大量的goroutine抢锁可能会导致CPU无意义的资源浪费。有一个专有名词用来描述这种抢锁的场景：活锁。

活锁指的是程序看起来在正常执行，但实际上CPU周期被浪费在抢锁，而非执行任务上，从而程序整体的执行效率低下。活锁的问题定位起来要麻烦很多。
所以在单机场景下，不建议使用这种锁。

*/

// Lock try lock
type Lock struct {
	c chan struct{}
}

// NewLock generate a try lock
func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}

// Lock try lock, return lock result
func (l Lock) Lock() bool {
	lockResult := false
	select {
	case <-l.c:
		lockResult = true
	default:
	}
	return lockResult
}

// Unlock , Unlock the try lock
func (l Lock) Unlock() {
	l.c <- struct{}{}
}

func TestTryLock(t *testing.T) {
	var l = NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.Lock() {
				// log error
				println("lock failed")
				return
			}
			counter++
			println("current counter", counter)
			l.Unlock()
		}()
	}
	wg.Wait()
}
