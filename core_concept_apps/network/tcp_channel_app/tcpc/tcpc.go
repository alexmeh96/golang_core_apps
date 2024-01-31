package tcpc

import (
	"encoding/gob"
	"log"
	"net"
	"sync"
	"time"
)

type TCPC[T any] struct {
	listenAddr string
	remoteAddr string

	SendChan     chan T
	RecvChan     chan T
	outboundConn net.Conn
	ln           net.Listener
	wg           sync.WaitGroup
}

func New[T any](listenAddr, remoteAddr string) (*TCPC[T], error) {
	tcpc := &TCPC[T]{
		listenAddr: listenAddr,
		remoteAddr: remoteAddr,
		SendChan:   make(chan T, 10),
		RecvChan:   make(chan T, 10),
	}

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return nil, err
	}
	tcpc.ln = ln

	go tcpc.loop()
	go tcpc.acceptLoop()
	go tcpc.dialRemoteAndRead()

	return tcpc, nil
}

func (t *TCPC[T]) loop() {
	t.wg.Wait()

	for {
		msg := <-t.SendChan
		if err := gob.NewEncoder(t.outboundConn).Encode(&msg); err != nil {
			log.Println(err)
			return
		}
		log.Printf("sended msg over the wire to %s: %v", t.remoteAddr, msg)
	}
}

func (t *TCPC[T]) acceptLoop() {
	defer func() {
		t.ln.Close()
	}()

	for {
		conn, err := t.ln.Accept()
		if err != nil {
			log.Println("accept error:", err)
			return
		}

		log.Printf("sender connected %s", conn.RemoteAddr())

		go t.handleConn(conn)
	}
}

func (t *TCPC[T]) handleConn(conn net.Conn) {
	defer func() {
		conn.Close()
	}()

	for {
		var msg T
		if err := gob.NewDecoder(conn).Decode(&msg); err != nil {
			log.Println(err)
			return
		}
		t.RecvChan <- msg
	}
}

func (t *TCPC[T]) dialRemoteAndRead() {
	t.wg.Add(1)

	for {
		conn, err := net.Dial("tcp", t.remoteAddr)
		if err != nil {
			log.Printf("dial error (%s)", err)
			time.Sleep(time.Second * 3)
		} else {
			t.outboundConn = conn
			break
		}
	}

	t.wg.Done()
}
