package main

import (
	"net/http"
)

// 数据爬虫，去重，写入kafka
func main() {
	http.HandleFunc("/crawlNews", func(w http.ResponseWriter, r *http.Request) {
		redisClient := newRedisClient()
		wangYiXinWenNews := removeDuplicate(redisClient, wangYiXinWenDataGet())
		tianXingTouTiaoNews := removeDuplicate(redisClient, tianXingTouTiaoDataGet())
		deliverMessageToKafka("wangYiXinWenNews", wangYiXinWenNews)
		deliverMessageToKafka("tianXingTouTiaoNews", tianXingTouTiaoNews)
	})
	http.ListenAndServe(":80", nil)

}


