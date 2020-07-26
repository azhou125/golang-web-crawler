package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"context"
	"github.com/go-redis/redis"
)

////////数据归一化///////

// 网易新闻struct
type WangYiXinWenRaw struct {
	BBM54PGAwangning []struct {
		LiveInfo interface{} `json:"liveInfo"`
		Docid string `json:"docid"`
		Source string `json:"source"`
		Title string `json:"title"`
		Priority int `json:"priority"`
		HasImg int `json:"hasImg"`
		URL string `json:"url"`
		CommentCount int `json:"commentCount"`
		Imgsrc3Gtype string `json:"imgsrc3gtype"`
		Stitle string `json:"stitle"`
		Digest string `json:"digest"`
		Imgsrc string `json:"imgsrc"`
		Ptime string `json:"ptime"`
	} `json:"BBM54PGAwangning"`
}

// 天行头条struct
type TianXingTouTiaoRaw struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Newslist []struct {
		Ctime string `json:"ctime"`
		Title string `json:"title"`
		Description string `json:"description"`
		PicURL string `json:"picUrl"`
		URL string `json:"url"`
		Source string `json:"source"`
	} `json:"newslist"`
}

func Json2Go(body []byte, rawNews interface{}) {
	err := json.Unmarshal([]byte(body), &rawNews)
	if err != nil {
		fmt.Println("Couldn't convert json to go", err)
	}
	//fmt.Println("AAA: ",rawNews)
}
func Go2Json(data interface{}) []byte{
	stdJson, err := json.Marshal(data)
	if err !=nil {
		fmt.Println("Couldn't convert go objects back to json", err)
	}
	//fmt.Println(fmt.Sprintf(string(stdJson)))
	return stdJson
}

// 标准struct
type StdNew struct{
	Timestamp string `json:"timestamp"`
	Source string `json:"source"`
	Title string `json:"title"`
	Body string `json:"body"`
	URL string `json:"URL"`
	//Types []string `json:"types"`
}
func TianXingTouTiaoGO2Std(txtt TianXingTouTiaoRaw) []StdNew {
	ret := []StdNew{}
	for i:= range txtt.Newslist{
		var item StdNew
		item.Source = txtt.Newslist[i].Source
		item.Timestamp = txtt.Newslist[i].Ctime
		item.Title = txtt.Newslist[i].Title
		item.Body = txtt.Newslist[i].Description
		item.URL = txtt.Newslist[i].URL
		ret = append(ret, item)
	}
	return ret
}
func WangYiXinWenGO2StdGo(wyxw WangYiXinWenRaw) []StdNew {
	ret := []StdNew{}
	for i:= range wyxw.BBM54PGAwangning{
		var item StdNew
		item.Source = wyxw.BBM54PGAwangning[i].Source
		item.Timestamp = wyxw.BBM54PGAwangning[i].Ptime
		item.Title = wyxw.BBM54PGAwangning[i].Title
		item.Body = wyxw.BBM54PGAwangning[i].Digest
		item.URL = wyxw.BBM54PGAwangning[i].URL
		ret = append(ret, item)
	}
	return ret
}
//func ZHGO2Std(zh ZongHeNewsRaw) []StdNew {
//	ret := []StdNew{}
//	for i:= range zh.Newslist{
//		var item StdNew
//		item.Source = "zonghe"
//		item.Timestamp = time.Now().Format("2006-01-02 15:04:05")
//		item.Title = zh.Newslist[i].Title
//		item.Body = zh.Newslist[i].Description
//		item.URL = zh.Newslist[i].URL
//		ret = append(ret, item)
//	}
//	return ret
//}

/////调用第三方API/////
func wangYiXinWenDataGet() []byte {
	url := "https://3g.163.com/touch/reconstruct/article/list/BBM54PGAwangning/0-10.html"

	resp, err := http.Get(url)
	if err !=nil {
		fmt.Println("Couldn't fetch news", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err !=nil {
		fmt.Println("Couldn't read news", err)
	}

	//fmt.Println(string(body))

	data := WangYiXinWenRaw{}
	Json2Go([]byte(body[9: len(body)-1]), &data)

	//fmt.Println(data)

	stdData := WangYiXinWenGO2StdGo(data)

	stdJson := Go2Json(stdData)
	fmt.Println(string(stdJson))

	return stdJson
}

//func tianXingTouTiaoDataGet() TianXingTouTiaoDataList{
//	url := "http://api.tianapi.com/topnews/index?key=4703f453980fd2a17b6413701b591c4b"
//
//	resp, _ := http.Get(url)
//	body, _ := ioutil.ReadAll(resp.Body)
//	resp.Body.Close()
//
//	s := TianXingTouTiaoDataList{}
//	json.Unmarshal([]byte(body), &s)
//	//fmt.Println(fmt.Sprintf("%+v",s))
//	//fmt.Println()
//	return s
//}

/////提取正文/////
//func getArticleContent(value GeneralDataList) {
//	panic("not implemented")
//}

/////调用kafka客户端发送消息//////
func deliverMessageToKafka(topic string, data []byte){

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	Produce(p,topic,nil,data)
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
func receiveMessageFromKafka(topic string) {

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
		receivedNews := StdNew{}
		Json2Go(msg.Value,&receivedNews)
	}


}





var ctx = context.Background()
func ExampleNewClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

/////////
func main() {
	//deliverMessageToKafka("topic1",wangYiXinWenDataGet())
	//receiveMessageFromKafka("topic1")
	ExampleClient()



}
