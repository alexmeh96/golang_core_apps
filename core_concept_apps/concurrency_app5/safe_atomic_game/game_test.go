package safe_atomic_game

import "testing"

func TestGame(t *testing.T) {
	player := NewPlayer()
	go startUILoop(player)
	startGameLoop(player)
}
