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

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(data),
	}

	p.Produce(msg,nil)

	// Wait for message deliveries before shutting down
	p.Flush(5 * 1000)
}



/////////
func main() {
	//wangYiXinWenDataGet()
	//tianXingTouTiaoDataGet()

	deliverMessageToKafka("topic1",wangYiXinWenDataGet())
	deliverMessageToKafka("topic2",tianXingTouTiaoDataGet())
}
