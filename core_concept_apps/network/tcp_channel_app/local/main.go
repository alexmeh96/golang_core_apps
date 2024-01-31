package main

import (
	"fmt"
	"log"
	"tcp_channel_app/tcpc"
)

func main() {
	channel, err := tcpc.New[string](":3000", ":4000")
	if err != nil {
		log.Fatal(err)
	}

	channel.SendChan <- "Hello"

	msg := <-channel.RecvChan

	fmt.Println("received msg from channel (:4000) over TCP: ", msg)

	select {}
}
