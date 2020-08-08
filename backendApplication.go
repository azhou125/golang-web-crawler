package main

func main() {

	redisClient := newRedisClient()
	uniqueStdNews1 := removeDuplicate(redisClient, wangYiXinWenDataGet())
	deliverMessageToKafka("topic1",uniqueStdNews1)
	receiveMessageFromKafka("topic1")
	//uniqueStdNews2 := removeDuplicate(redisClient, tianXingTouTiaoDataGet())
	//deliverMessageToKafka("topic2",uniqueStdNews2)
	//receiveMessageFromKafka("topic2")

}
