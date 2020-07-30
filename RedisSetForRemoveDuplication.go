package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"reflect"
)

var ctx = context.Background()

func newRedisClient()  *redis.Client{
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println("type:", reflect.TypeOf(rdb))
	return rdb
}

func removeDuplicate(rdb *redis.Client, stdNews []StdNew) []StdNew{
	ret := []StdNew{}
	for i:= range stdNews {
		key := stdNews[i].Title
		_, err := rdb.Get(ctx, key).Result()
		if err == redis.Nil {
			fmt.Println("This news is never seen before~~~~")
			err := rdb.Set(ctx, key, nil, 0).Err()
			if err != nil {
				panic(err)
			}
			ret = append(ret, stdNews[i])
		} else if err != nil {
			panic(err)
		} else {
			fmt.Println("Duplicate News Found! Title: "+key)
		}
	}
	return ret
}




//func ExampleNewClient() {
//	rdb := redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379",
//		Password: "", // no password set
//		DB:       0,  // use default DB
//	})
//
//	pong, err := rdb.Ping(ctx).Result()
//	fmt.Println(pong, err)
//	// Output: PONG <nil>
//}

//func ExampleClient() {
//	rdb := redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379",
//		Password: "", // no password set
//		DB:       0,  // use default DB
//	})
//	err := rdb.Set(ctx, "key", "value", 0).Err()
//	if err != nil {
//		panic(err)
//	}
//
//	val, err := rdb.Get(ctx, "key").Result()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("key", val)
//
//	val2, err := rdb.Get(ctx, "key2").Result()
//	if err == redis.Nil {
//		fmt.Println("key2 does not exist")
//	} else if err != nil {
//		panic(err)
//	} else {
//		fmt.Println("key2", val2)
//	}
//	// Output: key value
//	// key2 does not exist
//}
