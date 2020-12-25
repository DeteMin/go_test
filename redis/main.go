package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var (
	redisClient *redis.Client
)

func Init(){
	client := redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	redisClient = client
}


func main(){
	Init()

	//err := SetKey("way","go-redis", 10 * time.Minute)
	//if err != nil {
	//	fmt.Sprintf("err is:%v", err)
	//	return
	//}
	//
	//value, err := GetKey("way")
	//if err != nil {
	//	fmt.Println(fmt.Sprintf("err is:%v", err))
	//	return
	//}
	//fmt.Println(fmt.Sprintf("way is:%v", value))

	//err := HSET("SHARE-REWARD", "order_id", "1001")
	//if err != nil {
	//	fmt.Sprintf("err is:%v", err)
	//	return
	//}
	//resultMap, err := redisClient.HGetAll("SHARE-REWARD").Result()
	//if err != nil {
	//	fmt.Sprintf("err is:%v", err)
	//	return
	//}
	//for k, v := range resultMap {
	//	fmt.Println(fmt.Sprintf("key:%v  value:%v", k, v))
	//}

	doctor := Doctor{2, "钟南山2", 83, 1, time.Now()}
	doctorJson, _ := json.Marshal(doctor)
	err := SetList("order-list", doctorJson)
	if err != nil {
		fmt.Sprintf("err is:%v", err)
		return
	}
	list, _ := redisClient.SMembers("order-list").Result()
	for _, v := range list {
		var doctor2 Doctor
		//反序列化
		json.Unmarshal([]byte(v), &doctor2)
		fmt.Println("doctor2", doctor2)
	}

}

//存普通key
func SetKey(key, value string, expiration time.Duration) error {
	err := redisClient.Set(key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

//取普通key
func GetKey(key string) (string, error) {
	value, err := redisClient.Get(key).Result()
	if err != nil {
		return "", err
	}
	return value,nil
}

func DelKey(key string) error {
	_, err := redisClient.Del(key).Result()
	if err != nil {
		return err
	}
	return nil
}

type Doctor struct {
	ID      int64
	Name    string
	Age     int
	Sex     int
	AddTime time.Time
}

func HSET(key, field, value string) error {
	err := redisClient.HSet(key, field, value).Err()
	if err != nil {
		return err
	}
	err = redisClient.Expire(key, 3 * time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

func SetList(key string, value []byte) error {
	redisClient.SAdd(key, value)

	doctor := Doctor{1, "钟南山", 83, 1, time.Now()}
	doctorJson, _ := json.Marshal(doctor)
	err := redisClient.SAdd(key, doctorJson).Err()
	if err != nil {
		return err
	}

	redisClient.HDel(key,"1001").Result()

	return nil
}