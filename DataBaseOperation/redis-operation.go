package DataBaseOperation

import (
	"github.com/go-redis/redis"
)

type JsonData struct {
	Header  string      `json:"header"`
	Payload interface{} `json:"payload"`
}

var jd JsonData

func RunnerRedis() {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379",
		Password: "",
		DB:       0})
	TrunkOne(client)
}

func TrunkOne(client *redis.Client) {

	strSlice:=client.LRange("newusers",0,50)
	strSlice.Result()
}
