package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func subscriber(client *redis.Client) {
	pubsub := client.Subscribe("mychannel")
	defer pubsub.Close()

	// 处理订阅接收到的消息
	for {
		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			return
		}

		fmt.Println(msg.Channel, msg.Payload)
	}
}

func publisher(client *redis.Client) {
	for {
		// 发布消息到频道
		err := client.Publish("mychannel", "hello").Err()
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	go subscriber(client)
	go publisher(client)

	<-make(chan struct{})
}
