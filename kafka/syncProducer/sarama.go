package syncProducer

import "fmt"
import "github.com/Shopify/sarama"


func BrokerInfo() {
	broker := sarama.NewBroker(broker)
	err := broker.Open(nil)
	if err != nil {
		panic(err)
	}

	request :=  sarama.MetadataRequest{Topics: []string{topic}}
	response, err := broker.GetMetadata(&request)
	if err != nil {
		_ = broker.Close()
		panic(err)
	}

	fmt.Println("There are", len(response.Topics), "topics active in the cluster.")

	if err = broker.Close(); err != nil {
		panic(err)
	}
}
