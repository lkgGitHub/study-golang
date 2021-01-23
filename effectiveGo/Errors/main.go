package main

import (
	"os"
	"syscall"
)

func main() {
	filename := "aaa"
	for try := 0; try < 2; try++ {
		_, err := os.Create(filename)
		if err == nil {
			return
		}
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
			println("deleteTempFiles()")
			continue
		}
		return
	}
}

//// PathError 记录一个错误以及产生该错误的路径和操作。
//type PathError struct {
//	Op   string // "open"、"unlink" 等等。
//	Path string // 相关联的文件。
//	Err  error  // 由系统调用返回。
//}
//
//func (e *PathError) Error() string {
//	return e.Op + " " + e.Path + ": " + e.Err.Error()
//}
