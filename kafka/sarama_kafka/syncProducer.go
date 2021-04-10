package saramaKafka

import (
	"log"

	"github.com/Shopify/sarama"
)

func SyncProducer() {
	producer, err := sarama.NewSyncProducer([]string{broker}, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	s := `{"DeviceInfo":{"LogicDeviceID":"3c2c5554-a694-4987-b450-275e4e1480c8","PhysicalDeviceID":"1428007134","DeviceGroupID":"164ab196-b471-4ce6-82bb-943e25343917","TenantID":"09e95c6c-dcdf-40c8-94e8-bdc2fb462764","ClusterRepoID":"b8031b33-6287-4d9a-89f4-fee5744a59f9","SceneType":"other","ClusterVersion":1,"RegisterReposInfo":[{"ID":"02af20a0-38db-432e-8e2e-fce67f7bb751","Precise":0,"RankRegister":false}],"CaptureRepoInfo":{"face":{"repo":"","rankRegister":false},"humanBodyFeature":{"repo":"","rankRegister":false},"motorVehicleFeature":{"repo":"","rankRegister":false},"nonMotorVehicleFeature":{"repo":"","rankRegister":false},"entranceFace":{"repo":"","rankRegister":false}}},"AiRequest":{"AsyncRequest":true,"DeviceID":"1428007134","Objects":[{"ObjectID":"83fb21d5-0323-4d1f-9e5d-8902e5221451","ObjectType":"face","CaptureTime":1616073031587,"Image":{"URL":"http://dgtest.cn-bj.ufileos.com/13cd1a6b-0bec-40e8-a184-7d20c5684d6a_14f7091a-976f-4476-9a5a-247898d562dc","BinData":"","Width":272,"Height":200,"X":360,"Y":0},"OriginObjectID":""},{"ObjectID":"f13fcde2-075a-4e25-9f78-3f7aaf80fe8c","ObjectType":"scenePicture","CaptureTime":1616073031587,"Image":{"URL":"http://bptest.cn-bj.ufileos.com/13cd1a6b-0bec-40e8-a184-7d20c5684d6a","BinData":"","Width":1920,"Height":1080,"X":0,"Y":0},"OriginObjectID":""}],"Relationships":[{"ObjectID1":"f13fcde2-075a-4e25-9f78-3f7aaf80fe8c","ObjectID2":"83fb21d5-0323-4d1f-9e5d-8902e5221451","Relation":"scene_of"}],"Extra":{"Cached":"","Extend":"","Temperature":"","Trace":""},"IsUniqueObjectID":true}}`
	msg := &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(s)}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}
}
