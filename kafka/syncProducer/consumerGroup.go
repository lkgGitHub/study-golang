package syncProducer

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
)

type ConsumerGroupHandler struct{}

func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (c ConsumerGroupHandler) ConsumeClaim(cgs sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		cgs.MarkMessage(msg, "")
	}
	return nil
}

func ConsumerGroup() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	group, err := sarama.NewConsumerGroup([]string{broker}, topic, config)
	if err != nil {
		panic(err)
	}
	defer func() { _ = group.Close() }()

	// Track errors
	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR:", err.Error())
		}
	}()

	// Iterate over consumer sessions.
	ctx := context.Background()
	topics := []string{topic}
	handler := ConsumerGroupHandler{}
	for {
		// `Consume` should be called inside an infinite loop, when a
		// server-side reBalance happens, the consumer session will need to be
		// recreated to get the new claims
		err := group.Consume(ctx, topics, handler)
		if err != nil {
			panic(err)
		}
	}

}
