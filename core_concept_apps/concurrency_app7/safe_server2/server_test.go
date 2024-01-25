package safe_server2

import (
	"fmt"
	"testing"
)

func TestAddPlayer(t *testing.T) {
	server := NewServer()

	for i := 0; i < 10; i++ {
		player := &Player{
			Name: fmt.Sprintf("player_%d", i),
		}
		go server.handleNewPlayer(player)
	}
}

func TestSetFoo(t *testing.T) {
	server := NewServer()

	for i := 0; i < 10; i++ {
		go server.handleSetFoo(i)
	}
}
