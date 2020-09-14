package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	// 1. 初始化参数
	root :=  flag.String("f", ".", "files")
	verbose := flag.Bool("v", false, "show verbose progress messages")
	flag.Parse()
	roots := []string{*root, "/Users/lkg/go", "/Users/lkg/software", "/Users/lkg/Downloads"}
	//if len(roots) == 0 {
	//	roots = []string{"."}
	//}

	// 2. 遍历文件
	fmt.Println("root:", roots)
	fileSeizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots{
		n.Add(1)
		go walkDir(root, &n, fileSeizes)
	}
	go func() {
		n.Wait()
		close(fileSeizes)
	}()
	//go func() {
	//	for _, root := range roots{
	//		walkDir(root, fileSeizes)
	//	}
	//	close(fileSeizes)
	//}()

	// 3. 计算文件数和文件大小之和
	var tick <- chan time.Time
	if *verbose {
		tick = time.Tick(500*time.Millisecond)
	}
	// 打印中间过程
	var nfiles, nbytes int64
	// 这里的break语句用到了标签break("loop"），这样可以同时终结select和for两个循环；
	// 如果没有用标签就break的话只会退出内层的select循环，而外层的for循环会使之进入下一轮select循环。”
loop:
	for {
		select {
		case size, ok := <-fileSeizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <- tick:
			printDiskUsage(nfiles, nbytes)
		}

	}
	// 只打印最后结果
	//for size:= range fileSeizes{
	//	nfiles++
	//	nbytes += size
	//}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles int64, nbytes int64) {
	fmt.Printf("%d files  %.2f GB\n", nfiles, float64(nbytes)/1e9) // “1E9”，意思是 1 *（10^9)
}

func walkDir(dir string, s *sync.WaitGroup, fileSize chan<- int64) {
	defer s.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir){
		if entry.IsDir(){
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, nil, fileSize)
			time.Sleep(1*time.Millisecond) // 加个等待，否则太快了
		}else {
			fileSize <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)
// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}:
	case <- done:
		return nil
	}
	defer func() {<-sema}() // 控制并发，保并发只有20个

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <- done:
		return true
	default:
		return false
	}
}

func direntsPrint()  {
	dir := "/Users/lkg/Downloads"
	entries, err := ioutil.ReadDir(dir)
	if err != nil{
		fmt.Fprintf(os.Stderr, "du1: %v \n", err)
	}
	for i, e := range entries {
		println(i, "e: ", e.Size(), ": ", e.Name())
	}
}
