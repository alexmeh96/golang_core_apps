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
	msgCh  chan Message
	quitCh chan struct{}
}

func (s *Server) StartAndListen() {

loop:
	for {
		select {
		// block here until someone is sending a message to the channel
		case msg := <-s.msgCh:
			fmt.Printf("received message from: %s payload %s\n", msg.From, msg.Payload)
		case <-s.quitCh:
			fmt.Println("the server is doing a gracefull shutdown")
			// logic for the gracefull down
			break loop
		}
	}

	fmt.Println("the server is shut down!")
}

func sendMessageToServer(msgCh chan Message, payload string) {
	msg := Message{
		From:    "Alex",
		Payload: payload,
	}

	msgCh <- msg
}

func graceFullQuitServer(quitCh chan struct{}) {
	close(quitCh)
}

func main() {
	s := &Server{
		msgCh:  make(chan Message),
		quitCh: make(chan struct{}),
	}

	go s.StartAndListen()

	go func() {
		time.Sleep(2 * time.Second)
		sendMessageToServer(s.msgCh, "Hello Den!")
	}()

	go func() {
		time.Sleep(4 * time.Second)
		graceFullQuitServer(s.quitCh)
	}()

	select {}
}
