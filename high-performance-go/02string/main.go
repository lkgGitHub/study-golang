package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
)

/*
结论：
+ 和 fmt.Sprintf 的效率是最低的，和其余的方式相比，性能相差约 1000 倍，而且消耗了超过 1000 倍的内存。
fmt.Sprintf 通常是用来格式化字符串
综合易用性和性能，一般推荐使用 strings.Builder 来拼接字符串。

1. 字符串最高效的拼接方式是结合预分配内存方式 Grow 使用 string.Builder
2. 当使用 + 拼接字符串时，生成新字符串，需要开辟新的空间
3. 当使用 strings.Builder，bytes.Buffer 或 []byte 的内存是按倍数申请的(前期)，在原基础上不断增加
4. strings.Builder 比 bytes.Buffer 性能更快，一个重要区别在于 bytes.Buffer 转化为字符串重新申请了一块空间存放生成的字符串变量；而 strings.Builder 直接将底层的 []byte 转换成字符串类型返回
*/

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// 1. 使用 +
func plusConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}

// 2. 使用 fmt.Sprintf
func sprintfConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", s, str)
	}
	return s
}

// 3. 使用 strings.Builder
func builderConcat(n int, str string) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

// 3.1 使用 strings.Builder，预分配内存的方式 Grow
func preBuilderConcat(n int, str string) string {
	var builder strings.Builder
	builder.Grow(n * len(str))
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

// 4. 使用 bytes.Buffer
func bufferConcat(n int, s string) string {
	buf := new(bytes.Buffer)
	for i := 0; i < n; i++ {
		buf.WriteString(s)
	}
	return buf.String()
}

// 5. 使用 []byte
func byteConcat(n int, str string) string {
	buf := make([]byte, 0)
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}

// 5.1 使用 []byte。如果长度是可预知的，那么创建 []byte 时，我们还可以预分配切片的容量(cap)。
func preByteConcat(n int, str string) string {
	buf := make([]byte, 0, n*len(str))
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}
