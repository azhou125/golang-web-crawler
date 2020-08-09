package SharedFiles

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/go-redis/redis/v8"
	"reflect"
)

var ctx = context.Background()

func NewRedisClient()  *redis.Client{
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println("type:", reflect.TypeOf(rdb))
	return rdb
}

func RemoveDuplicate(rdb *redis.Client, stdNews []StdNew) []StdNew{
	ret := []StdNew{}
	for i:= range stdNews {
		key := stdNews[i].Title
		fmt.Printf("%x", md5.Sum([]byte(key)))
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