package main

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

/////调用kafka客户端发送消息//////
func deliverMessageToKafka(topic string, stdNews []StdNew){
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	for i:= range stdNews {
		data := Go2Json(stdNews[i])
		Produce(p,topic,nil,data)
	}

	//p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	//if err != nil {
	//	panic(err)
	//}
	//defer p.Close()
	//
	//Produce(p,topic,nil,data)
}

func Produce(producer *kafka.Producer, topic string, key, data []byte) error {
	deliveryChan := make(chan kafka.Event)
	defer close(deliveryChan)
	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
		Key:            key,
	}, deliveryChan)
	if err != nil {
		return nil
	}
	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivery message to topic %s [%d] at offset %v\n",
			m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		fmt.Println()
		fmt.Println()
	}
	return nil
}


/////从kafka客户端接收消息//////
func receiveMessageFromKafka(topic string){

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{topic, "^aRegex.*[Tt]opic"}, nil)

	esclient := getESClient()

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}

		receivedNews := StdNew{}
		Json2Go([]byte(msg.Value),&receivedNews)

		insertNews(esclient, receivedNews)
	}
}