package singleton

import (
	"fmt"
	"sync"
)

// 全局实例者
type singleton struct {
	data int
}

// 定义一个包级别的private实例变量
var sin *singleton

// 同步Once,保证每次调用时，只有第一次生效
var once sync.Once

func GetSingleton() *singleton {
	once.Do(func() {
		sin = &singleton{16}
	})
	fmt.Printf("实例对象的信息和地址,sin: %v, &sin: %v \n", sin, &sin)
	return sin
}

type singletonMap map[string]string

var sinMap singletonMap

func NewSingletonMap() singletonMap {
	once.Do(func() {
		sinMap = make(singletonMap)
	})
	return sinMap
}
