package safe_server2

import "fmt"

type Player struct {
	Name string
}

type SetFooMsg struct {
	value int
}

type GameState struct {
	players []*Player
	foo     int

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
	case *SetFooMsg:
		g.SetFoo(msg)
	default:
		panic("invalid message received")
	}
}

func (g *GameState) addPlayer(p *Player) {
	g.players = append(g.players, p)

	fmt.Println("adding player:", p.Name)
}

func (g *GameState) SetFoo(foo *SetFooMsg) {
	g.foo = foo.value

	fmt.Println("setting foo:", foo.value)
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

func (s *Server) handleSetFoo(val int) error {
	s.gameState.Receive(&SetFooMsg{value: val})
	return nil
}

func (s *Server) handleNewPlayer(player *Player) error {
	s.gameState.Receive(player)
	return nil
}
