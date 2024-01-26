package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	sub := client.Subscribe(ctx, "coords")
	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", msg.Payload)
	}
}
