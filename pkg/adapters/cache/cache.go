package cache

import (
	"encoding/json"
	"log"
	"time"

	"95.85.87.178/BonoboGitServer/IQ-Services-backend.git/pkg/utils"
	"github.com/go-redis/redis"
	jsoniter "github.com/json-iterator/go"
)

var redis_client *redis.Client

func SetRedisClient(client *redis.Client) {
	redis_client = client
}

func Sub(redis_channel string) *redis.PubSub {
	return redis_client.Subscribe(redis_channel)
}

func PublishToRedisChannel(channel string, message interface{}) {
	log.Printf("publish to redis %s, message : %v\n", channel, message)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	m, err := json.Marshal(message)
	utils.CatchErr(err)
	redis_client.Publish(channel, m)
}

func SetToCash(id string, data interface{}, time time.Duration) {
	log.Printf("redis_client.Info() %v", redis_client.Info())
	result, err := json.Marshal(data)
	utils.CatchErr(err)
	cmd := redis_client.Set(id, result, time)
	log.Printf(" commant to redis fmt:String() %v\n\n\n", cmd.String())
}

func GetCash(id string) interface{} {
	var data interface{}
	cmd := redis_client.Get(id)
	byte_data, _ := cmd.Bytes()
	json.Unmarshal(byte_data, &data)
	return data
}
