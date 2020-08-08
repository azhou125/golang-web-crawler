package main
//
//import (
//	"fmt"
//	"net/http"
//)
//
//// 数据爬虫，去重，写入kafka
//func main() {
//	http.HandleFunc("/crawlNews", func(w http.ResponseWriter, r *http.Request) {
//		redisClient := newRedisClient()
//		wangYiXinWenNews := removeDuplicate(redisClient, wangYiXinWenDataGet())
//		tianXingTouTiaoNews := removeDuplicate(redisClient, tianXingTouTiaoDataGet())
//		deliverMessageToKafka("wangYiXinWenNews", wangYiXinWenNews)
//		deliverMessageToKafka("tianXingTouTiaoNews", tianXingTouTiaoNews)
//	})
//	http.ListenAndServe(":80", nil)
//
//}
//
//// 从Kafka读取消息，并写入Elastic Search
//func main1() {
//	cnt := 0
//	go func(){
//		receiveMessageFromKafka("topic1")
//		receiveMessageFromKafka("topic2")
//		cnt++
//	}()
//
//	http.HandleFunc("/isactive", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, `{"storeCount":%d}`, cnt)
//	})
//	http.ListenAndServe(":81", nil)
//}
//
//// RestApi的访问服务器
//func main2() {
//	startServer()
//}
//
//
