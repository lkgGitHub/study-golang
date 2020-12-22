package main

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

var mu sync.Mutex
var chain string

// fatal error: all goroutines are asleep - deadlock!
func TestDemo(t *testing.T) {
	chain = "main"
	A()
	fmt.Println(chain)
}
func A() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> A"
	B()
}
func B() {
	chain = chain + " --> B"
	C()
}
func C() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> C"
}

var rwMu sync.RWMutex
var count int

// fatal error: all goroutines are asleep - deadlock!
// RLock 读锁上加 写Lock 不被允许
func TestDemo2(t *testing.T) {
	go A2()
	time.Sleep(2 * time.Second)
	rwMu.Lock()
	defer rwMu.Unlock()
	count++
	fmt.Println(count)
}
func A2() {
	rwMu.RLock()
	defer rwMu.RUnlock()
	B2()
}
func B2() {
	time.Sleep(5 * time.Second)
	C2()
}
func C2() {
	rwMu.RLock()
	defer rwMu.RUnlock()
}

// panic: sync: WaitGroup is reused before previous Wait has returned
func TestDemo3(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		println("2. wg done")
		wg.Add(1)
	}()
	println("1. wg wait")
	wg.Wait()
}

// 4 双检查实现单例
func TestDemo4(t *testing.T) {
	o := &Once{}
	o.Do(func() {
		println("func")
	})
}

type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if o.done == 1 {
		return
	}

	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		o.done = 1
		f()
	}
}

type MyMutex struct {
	count int
	// A Mutex must not be copied after first use.
	sync.Mutex
}

// 变量声明将锁定值复制到 'mu2': 类型 'MyMutex' 为 'sync.Locker'
// 检查信息: Checks for locks mistakenly passed by value.
func TestDemo5(t *testing.T) {
	var mu MyMutex
	mu.Lock()
	var mu2 = mu // 变量声明将锁定值复制到 'mu2': 类型 'MyMutex' 为 'sync.Locker'
	println("over: mu2 = mu")
	mu.count++
	mu.Unlock()
	println("mu2.Lock")
	mu2.Lock()
	println("不会走")
	mu2.count++
	mu2.Unlock()
	fmt.Println(mu.count, mu2.count)
}

var pool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

func TestDemo6(t *testing.T) {
	go func() {
		for {
			processRequest(1 << 28) // 256MiB
		}
	}()
	for i := 0; i < 1000; i++ {
		go func() {
			for {
				processRequest(1 << 10) // 1KiB
			}
		}()
	}
	var stats runtime.MemStats
	for i := 0; ; i++ {
		runtime.ReadMemStats(&stats)
		fmt.Printf("Cycle %d: %dB\n", i, stats.Alloc)
		time.Sleep(time.Second)
		runtime.GC()
	}
}
func processRequest(size int) {
	b := pool.Get().(*bytes.Buffer)
	time.Sleep(500 * time.Millisecond)
	b.Grow(size)
	pool.Put(b)
	time.Sleep(1 * time.Millisecond)
}
