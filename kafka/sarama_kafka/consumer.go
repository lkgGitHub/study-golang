package saramaKafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
)

// todo 1. 制定了分区，只能消费单个分区
// todo 2. 消费完没有通知kafka，已经消费。重复消费
func Consumer() {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{broker}, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}
}
