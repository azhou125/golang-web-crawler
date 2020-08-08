package main
//
//import (
//	"fmt"
//	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
//)
//
///////调用kafka客户端发送消息//////
//func deliverMessageToKafka(topic string, stdNews []StdNew){
//	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
//	if err != nil {
//		panic(err)
//	}
//	defer p.Close()
//
//	for i:= range stdNews {
//		data := Go2Json(stdNews[i])
//		Produce(p,topic,nil,data)
//	}
//}
//
//func Produce(producer *kafka.Producer, topic string, key, data []byte) error {
//	deliveryChan := make(chan kafka.Event)
//	defer close(deliveryChan)
//	err := producer.Produce(&kafka.Message{
//		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
//		Value:          data,
//		Key:            key,
//	}, deliveryChan)
//	if err != nil {
//		return nil
//	}
//	e := <-deliveryChan
//	m := e.(*kafka.Message)
//	if m.TopicPartition.Error != nil {
//		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
//	} else {
//		fmt.Printf("Delivery message to topic %s [%d] at offset %v\n",
//			m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
//		fmt.Println()
//		fmt.Println()
//	}
//	return nil
//}
