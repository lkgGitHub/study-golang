package syncProducer

import "testing"

func TestBrokerInfo(t *testing.T) {
	BrokerInfo()
}

func TestSyncProducer(t *testing.T) {
	SyncProducer()
}

func TestAsyncProducerGoroutines(t *testing.T) {
	AsyncProducerGoroutines()
}

func TestAsyncProducerSelect(t *testing.T) {
	AsyncProducerSelect()
}