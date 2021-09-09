package concurrent

import (
	"sync"
	"time"
)

const cost = time.Microsecond

type RW interface {
	Write()
	Read()
}

type Lock struct {
	count int
	mu sync.Mutex
}

func (l *Lock) Write() {
	l.mu.Lock()
	l.count++
	time.Sleep(cost)
	l.mu.Unlock()
}

func (l *Lock) Read() {
	l.mu.Lock()
	_ = l.count
	time.Sleep(cost)
	l.mu.Unlock()
}

type RWLock struct {
	count int
	mu sync.RWMutex
}

func (r *RWLock) Write() {
	r.mu.Lock()
	r.count++
	time.Sleep(cost)
	r.mu.Unlock()
}

func (r *RWLock) Read() {
	r.mu.RLock()
	_ = r.count
	time.Sleep(cost)
	r.mu.RUnlock()
}
