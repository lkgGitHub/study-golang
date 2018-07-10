package main

import (
	"strings"
	"fmt"
	"time"
)

type LogProcess struct {
	rc chan string
	wc chan string
	path string
	influxDBDsn string
}

func (l *LogProcess) ReadFormFile() {
	line := "message"
	l.rc <- line
}

func (l *LogProcess) Process() {
	date := <- l.rc
	l.wc <- strings.ToUpper(date)
}

func (l *LogProcess) Write() {
	fmt.Println(<-l.wc)
}

func main() {
	lp := &LogProcess{
		rc: make(chan string),
		wc: make(chan string),
		path:"C:\\Users\\lkg\\Desktop\\TESO.txt",
		influxDBDsn:"C:\\Users\\lkg\\Desktop\\golangLog.txt",
	}
	go lp.ReadFormFile()
	go lp.Process()
	go lp.Write()

	time.Sleep(1*time.Second)
}