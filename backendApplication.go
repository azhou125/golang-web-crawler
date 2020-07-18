package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

)
////////数据归一化///////
type Data interface {

}
type WangYiXinWenData struct {
	data Data
	Timestamp uint `json:"ptime"`
	Source string `json:"source"`
	Title string `json:"title"`
	Url string `json:"url"`
	Desc string `json:"digest"`
}
type TianXingTouTiaoData struct {
	data Data
	Timestamp uint `json:"ctime"`
	Source string `json:"source"`
	Title string `json:"title"`
	Url string `json:"url"`
	Desc string `json:"description"`
}

type GeneralDataList interface{

}
type WangYiXinWenDataList struct {
	generalDataList GeneralDataList
	ArticleList []WangYiXinWenData `json:"BBM54PGAwangning"`
}
type TianXingTouTiaoDataList struct {
	generalDataList GeneralDataList
	ArticleList []TianXingTouTiaoData `json:"newslist"`
}

/////调用第三方API/////
func wangYiXinWenDataGet() GeneralDataList {
	url := "https://3g.163.com/touch/reconstruct/article/list/BBM54PGAwangning/0-10.html"

	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	s := WangYiXinWenDataList{}
	json.Unmarshal([]byte(body[9: len(body)-1]), &s)
	//fmt.Println(fmt.Sprintf("%+v",s))
	//fmt.Println()
	return s
}
func tianXingTouTiaoDataGet() GeneralDataList{
	url := "http://api.tianapi.com/topnews/index?key=4703f453980fd2a17b6413701b591c4b"

	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	s := TianXingTouTiaoDataList{}
	json.Unmarshal([]byte(body), &s)
	//fmt.Println(fmt.Sprintf("%+v",s))
	//fmt.Println()
	return s
}

/////提取正文/////
func getArticleContent(value GeneralDataList) {
	panic("not implemented")
}

/////调用kafka客户端发送消息//////
func deliverMessageToKafka(topic string, value GeneralDataList){

	data, _ := json.Marshal(value)

	fmt.Println(fmt.Sprintf(string(data)))
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "0.0.0.0:9092"})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	Produce(p,topic,nil,[]byte(data))

	//// Delivery report handler for produced messages
	//go func() {
	//	for e := range p.Events() {
	//		switch ev := e.(type) {
	//		case *kafka.Message:
	//			if ev.TopicPartition.Error != nil {
	//				fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
	//			} else {
	//				fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
	//			}
	//		}
	//	}
	//}()
	//
	//// Produce messages to topic
	//msg := &kafka.Message{
	//	TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	//	Value:          []byte(data),
	//}
	//
	//p.Produce(msg,nil)
	//
	//// Wait for message deliveries before shutting down
	//p.Flush(10 * 1000)
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

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

}




/////////
func main() {
	//wangYiXinWenDataGet()
	//tianXingTouTiaoDataGet()

	deliverMessageToKafka("topic1",wangYiXinWenDataGet())
	//deliverMessageToKafka("topic2",tianXingTouTiaoDataGet())
}
