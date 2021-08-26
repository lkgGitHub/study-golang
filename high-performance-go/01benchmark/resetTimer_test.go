package main

import (
	"testing"
	"time"
)

func BenchmarkFib2(b *testing.B) {
	time.Sleep(time.Second * 3) // 模拟耗时准备任务
	b.ResetTimer() // 重置定时器
	for n := 0; n < b.N; n++ {
		fib(30) // run fib(30) b.N times
	}
}


// bubbleSort 冒泡函数
func bubbleSort(nums []int) {
	for i:=0;i<len(nums);i++{
		for j:=1;j<len(nums) -1 ;j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] =  nums[j], nums[i]
			}
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums:= generateWithCap(10000)
		b.StartTimer()
		bubbleSort(nums)
	}
}

/*
每次函数调用前后需要一些准备工作和清理工作，我们可以使用 StopTimer 暂停计时以及使用 StartTimer 开始计时。

例如，如果测试一个冒泡函数的性能，每次调用冒泡函数前，需要随机生成一个数字序列，这是非常耗时的操作，这种场景下，就需要使用 StopTimer 和 StartTimer 避免将这部分时间计算在内。
*/