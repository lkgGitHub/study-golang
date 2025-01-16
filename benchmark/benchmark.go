package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type Metrics struct {
	SuccessCount        int64
	FailureCount        int64
	SuccessLatencyTotal time.Duration
	FailureLatencyTotal time.Duration
	RequestCount        int64
	TotalSuccessCount   int64
	TotalFailureCount   atomic.Int64
	mu                  sync.Mutex
}

func (m *Metrics) Record(success bool, latency time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if success {
		m.SuccessCount++
		m.TotalSuccessCount++
		m.SuccessLatencyTotal += latency
	} else {
		m.FailureCount++
		m.TotalFailureCount++
		m.FailureLatencyTotal += latency
	}
	m.RequestCount++
}

func (m *Metrics) Reset() (success, failure, totalSuccess, totalFailure int64, avgSuccessLatency, avgFailureLatency float64, qps int64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	success = m.SuccessCount
	failure = m.FailureCount
	totalSuccess = m.TotalSuccessCount
	totalFailure = m.TotalFailureCount
	totalRequests := m.RequestCount
	avgSuccessLatency = 0
	avgFailureLatency = 0
	if success > 0 {
		avgSuccessLatency = float64(m.SuccessLatencyTotal.Milliseconds()) / float64(success)
	}
	if failure > 0 {
		avgFailureLatency = float64(m.FailureLatencyTotal.Milliseconds()) / float64(failure)
	}
	qps = totalRequests

	// Reset counters for the next interval
	m.SuccessCount = 0
	m.FailureCount = 0
	m.RequestCount = 0
	m.SuccessLatencyTotal = 0
	m.FailureLatencyTotal = 0

	return
}

func worker(url string, metrics *Metrics, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		start := time.Now()
		resp, err := http.Get(url)
		latency := time.Since(start)

		if err != nil || resp.StatusCode != http.StatusOK {
			metrics.Record(false, latency)
			if resp != nil {
				resp.Body.Close()
			}
			continue
		}

		metrics.Record(true, latency)
		resp.Body.Close()
	}
}

func monitor(metrics *Metrics) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		success, failure, totalSuccess, totalFailure, avgSuccessLatency, avgFailureLatency, qps := metrics.Reset()
		fmt.Printf("QPS: %d, Success: %d, Failure: %d, Avg Success Latency: %.2f ms, Avg Failure Latency: %.2f ms, Total Success: %d, Total Failure: %d\n",
			qps, success, failure, avgSuccessLatency, avgFailureLatency, totalSuccess, totalFailure)
	}
}

func main() {
	url := "http://www.baidu.com" // Replace with the target URL
	numWorkers := 10              // Number of concurrent workers
	metrics := &Metrics{}

	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(url, metrics, &wg)
	}

	// Start monitoring
	go monitor(metrics)

	// Run for a fixed duration (e.g., 10 seconds)
	time.Sleep(10 * time.Second)

	fmt.Println("Stopping workers...")
	wg.Wait()
	fmt.Println("Test completed.")
}
