package unsafe_server

type Player struct {
	Name string
}

type GameState struct {
	players []*Player
}

func (g *GameState) addPlayer(p *Player) {
	g.players = append(g.players, p)
}

func NewGameState() *GameState {
	return &GameState{
		players: []*Player{},
	}
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
	s.gameState.addPlayer(player)
	return nil
}
