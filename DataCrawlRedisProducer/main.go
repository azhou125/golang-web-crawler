package main

import (
	"net/http"
	pkg "backendProject/SharedFiles"
)

// 数据爬虫，去重，写入kafka
func main() {
	http.HandleFunc("/crawlNews", func(w http.ResponseWriter, r *http.Request) {
		redisClient := pkg.NewRedisClient()
		wangYiXinWenNews := pkg.RemoveDuplicate(redisClient, pkg.WangYiXinWenDataGet())
		tianXingTouTiaoNews := pkg.RemoveDuplicate(redisClient, pkg.TianXingTouTiaoDataGet())
		pkg.DeliverMessageToKafka("wangYiXinWenNews", wangYiXinWenNews)
		pkg.DeliverMessageToKafka("tianXingTouTiaoNews", tianXingTouTiaoNews)
	})
	http.ListenAndServe(":80", nil)

}


