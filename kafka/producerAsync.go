package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

func ProducerAsync() {
	conf := sarama.NewConfig()
	//conf.Version = sarama.V0_10_2_0
	//conf.Producer.RequiredAcks = sarama.WaitForAll
	//conf.Producer.Return.Errors = true // 接收producer的error，下面细说用法
	conf.Producer.Compression = sarama.CompressionZSTD   // 压缩方式，如果kafka版本大于1.2，推荐使用zstd压缩
	conf.Producer.Flush.Messages = 10                    // 缓存条数
	conf.Producer.Flush.Frequency = 3 * time.Millisecond // 缓存时间

	asyncProducer, err := sarama.NewAsyncProducer([]string{Broken}, conf)
	if err != nil {
		panic(err.Error())
	}
	// 接受结束信号
	go func(p sarama.AsyncProducer) {
		success := p.Successes()
		for i := 0; ; i++ {
			select {
			case err := <-p.Errors():
				fmt.Printf("kafka producer send error %s \n", err.Err.Error())
			case producerMessage := <-success:
				log.Printf("%d > message sent to partition %d at offset %d \n", i, producerMessage.Partition, producerMessage.Offset)
			}
		}
	}(asyncProducer)

	for i := 0; ; i++ {
		time.Sleep(time.Millisecond * 500)

		value := "message"
		msg := &sarama.ProducerMessage{
			Topic:     Topic,
			Value:     sarama.StringEncoder(value),
			Timestamp: time.Now(),
		}
		asyncProducer.Input() <- msg
		log.Println(i)
	}
}
