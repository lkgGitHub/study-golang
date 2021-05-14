package main

import (
	"sync"
	"time"
)

/*
GO里面MAP如何实现key不存在 get操作等待 直到key存在或者超时，保证并发安全，且需要实现以下接口：
*/

type sp interface {
	//存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Out(key string, val interface{})
	//读取一个key，如果key不存在阻塞，等待key存在或者超时
	Rd(key string, timeout time.Duration) interface{}
}

type Map struct {
	c   map[string]*entry
	rwm *sync.RWMutex
}

type entry struct {
	ch      chan struct{}
	value   interface{}
	isExist bool
}

func (m Map) Out(key string, val interface{}) {
	m.rwm.Lock()
	defer m.rwm.Unlock()

	if e, ok := m.c[key]; ok {
		e.value = val
		e.isExist = true
		close(e.ch)
	} else {
		m.c[key] = &entry{ch: make(chan struct{}), value: val, isExist: true}
		close(e.ch)
	}
}

func (m Map) Rd(key string, timeout time.Duration) interface{} {
	m.rwm.RLock()
	defer m.rwm.RUnlock()

	v, ok := m.c[key]
	if ok && v.isExist {
		return v
	}
	if !ok {
		v = &entry{ch: make(chan struct{}), isExist: false}
		m.c[key] = v
		println("协程阻塞 -> ", key)
		select {
		case <-v.ch:
			return v.value
		case <-time.After(timeout):
			println("协程超时 -> ", key)
			return nil
		}
	} else {
		println("协程阻塞 -> ", key)
		select {
		case <-v.ch:
			return v.value
		case <-time.After(timeout):
			println("协程超时 -> ", key)
			return nil
		}
	}

}
