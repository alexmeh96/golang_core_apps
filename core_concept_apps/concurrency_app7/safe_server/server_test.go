package safe_server

import (
	"fmt"
	"testing"
)

func TestServer(t *testing.T) {
	server := NewServer()

	for i := 0; i < 10; i++ {
		player := &Player{
			Name: fmt.Sprintf("player_%d", i),
		}
		go server.handleNewPlayer(player)
	}
}
