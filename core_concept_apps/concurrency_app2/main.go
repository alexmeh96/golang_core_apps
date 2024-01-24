package main

import (
	"fmt"
	"time"
)

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgCh chan Message
}

func (s *Server) StartAndListen() {
	for {
		// block here until someone is sending a message to the channel
		msg := <-s.msgCh
		fmt.Printf("received message from: %s payload %s\n", msg.From, msg.Payload)
	}
}

func sendMessageToServer(msgCh chan Message, payload string) {
	msg := Message{
		From:    "Alex",
		Payload: payload,
	}

	msgCh <- msg
}

func main() {
	s := &Server{
		msgCh: make(chan Message),
	}

	go s.StartAndListen()

	sendMessageToServer(s.msgCh, "Hello Den!")
	time.Sleep(30 * time.Millisecond)
	sendMessageToServer(s.msgCh, "Hello Den2!")

	time.Sleep(80 * time.Millisecond)

}
