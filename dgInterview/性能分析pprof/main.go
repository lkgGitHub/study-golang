package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"sync"
	"time"
)

var cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		_ = pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	_ = initPprofMonitor()
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
		time.Sleep(time.Second * 3)
		runtime.GC()
	}
}

var pool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

func processRequest(size int) {
	b := pool.Get().(*bytes.Buffer)
	time.Sleep(500 * time.Millisecond)
	b.Grow(size)
	pool.Put(b)
	time.Sleep(1 * time.Millisecond)
}

func appendUnique(a []int, x int) []int {
	for _, y := range a {
		if x == y {
			return a
		}
	}
	return append(a, x)
}

func initPprofMonitor() error {
	pPort := 80

	var err error
	addr := ":" + strconv.Itoa(pPort)

	go func() {
		err = http.ListenAndServe(addr, nil)
		if err != nil {
			log.Printf("funcRetErr=http.ListenAndServe||err=%s", err.Error())
		}
	}()

	return err
}
