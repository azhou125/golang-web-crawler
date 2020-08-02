package main

func main() {

	uniqueStdNews := removeDuplicate(newRedisClient(), wangYiXinWenDataGet())
	deliverMessageToKafka("topic1",uniqueStdNews)
	receiveMessageFromKafka("topic1")

}
