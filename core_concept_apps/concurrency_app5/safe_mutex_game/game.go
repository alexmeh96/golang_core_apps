package safe_mutex_game

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Player struct {
	mu     sync.RWMutex
	health int
}

func (p *Player) getHealth() int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.health
}

func (p *Player) takeDamage(value int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.health -= value
}

func NewPlayer() *Player {
	return &Player{
		health: 100,
	}
}

func startUILoop(p *Player) {
	ticker := time.NewTicker(time.Second)

	for {
		fmt.Printf("player health: %d\n", p.getHealth())
		<-ticker.C
	}
}

func startGameLoop(p *Player) {
	ticker := time.NewTicker(time.Millisecond * 300)
	for {
		p.takeDamage(rand.Intn(40))
		if p.getHealth() <= 0 {
			fmt.Println("GAME OVER")
			break
		}
		<-ticker.C
	}
}
