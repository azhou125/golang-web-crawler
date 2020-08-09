package SharedFiles

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

/////从kafka客户端接收消息//////
func ReceiveMessageFromKafka(topic string){

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{topic, "^aRegex.*[Tt]opic"}, nil)

	esclient := GetESClient()

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

		InsertNews(esclient, receivedNews)
	}
}
