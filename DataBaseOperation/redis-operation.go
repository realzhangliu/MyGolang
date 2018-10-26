package DataBaseOperation

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"time"
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

	jd.Header = "h1"
	jd.Payload = "p1"
	data, _ := json.Marshal(&jd)

	client.Set("k1", data, time.Hour)

	res := client.Get("k1")
	res_data, err := res.Bytes()
	if err != nil {
		fmt.Println(err)
	}
	var jd2 JsonData
	json.Unmarshal(res_data,&jd2)
	fmt.Println(res_data)
	fmt.Println(client.Ping())
	defer client.Close()
}
