package baktu

// server contains info of all running games. This server only has 1
// instance and initialized 1 time
type server struct {
	games       map[string]*game
	destroyGame chan string
}

// newServer initialize server, this should be only called once
func newServer() *server {
	var s server
	s.games = make(map[string]*game)
	s.destroyGame = make(chan string)
	go func() {
		for {
			select {
			case gameID := <-s.destroyGame:
				delete(s.games, gameID)
			}
		}
	}()
	return &s
}

// cmdHandler universally interprets commands from any trigger
func (s *server) cmdHandler(userIn userInput) {
	gm, ok := s.games[userIn.gameID]
	switch userIn.command {
	case cmdUserJoin:
		// create new game if it doesn't exist
		if !ok {
			gm = s.newGame(userIn.gameID)
			s.games[userIn.gameID] = gm
			go gm.areaWaiting()
		}
		gm.cmd <- userIn
	case cmdUserHit:
		if ok {
			gm.cmd <- userIn
		}
	case cmdUserScore:
		gm.printScores()
	}
}
