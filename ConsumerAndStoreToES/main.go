package main

import (
	"fmt"
	"net/http"

	pkg "../SharedFiles"
)

// 从Kafka读取消息，并写入Elastic Search
func main() {
	cnt := 0
	go func(){
		pkg.ReceiveMessageFromKafka("topic1")
		pkg.ReceiveMessageFromKafka("topic2")
		cnt++
	}()

	http.HandleFunc("/isactive", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"storeCount":%d}`, cnt)
	})
	http.ListenAndServe(":81", nil)
}