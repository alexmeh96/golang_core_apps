package safe_server

import "fmt"

type Player struct {
	Name string
}

type GameState struct {
	players []*Player

	msgCh chan any
}

func (g *GameState) Receive(msg any) {
	g.msgCh <- msg
}

func (g *GameState) loop() {
	for msg := range g.msgCh {
		g.handleMessage(msg)
	}
}

func (g *GameState) handleMessage(message any) {
	switch msg := message.(type) {
	case *Player:
		g.addPlayer(msg)
	default:
		panic("invalid message received")
	}
}

func (g *GameState) addPlayer(p *Player) {
	g.players = append(g.players, p)

	fmt.Println("adding player:", p.Name)
}

func NewGameState() *GameState {
	g := &GameState{
		players: []*Player{},
		msgCh:   make(chan any, 10),
	}

	go g.loop()

	return g
}

type Server struct {
	gameState *GameState
}

func NewServer() *Server {
	return &Server{
		gameState: NewGameState(),
	}
}

func (s *Server) handleNewPlayer(player *Player) error {
	s.gameState.Receive(player)
	return nil
}
