package main

import (
	"fmt"
	"log"
	"tcp_channel_app/tcpc"
)

func main() {
	channel, err := tcpc.New[string](":4000", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	msg := <-channel.RecvChan

	fmt.Printf("received msg from channel (%s) over TCP: %s\n", ":3000", msg)

	channel.SendChan <- msg

	select {}
}
