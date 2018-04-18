package baktu

type handlerFunc func(*gameOutput)

// server contains info of all running games. This server only has 1
// instance and initialized 1 time.
type server struct {
	games     map[string]*game
	out       chan gameOutput
	responder *responder
}

// newServer initialize server, this should be only called once.
func newServer() *server {
	s := server{
		games:     make(map[string]*game),
		out:       make(chan gameOutput),
		responder: newResponder(),
	}
	go s.outputHandler()
	return &s
}

// inputHandler handles request from inputListener (which is api and/or
// webhook request). This instance universally interprets commands from
// any trigger.
func (s *server) inputHandler(userIn userInput) {
	gm, ok := s.games[userIn.gameID]
	switch userIn.command {
	case cmdUserJoin:
		// create new game if it doesn't exist
		if !ok {
			gm = s.newGame(userIn.gameID)
			s.games[userIn.gameID] = gm
			go gm.areaWaiting()
		}
		gm.in <- userIn
	case cmdUserHit:
		if ok {
			gm.in <- userIn
		}
	case cmdUserScore:
		gm.printScores()
	}
}

// outputHandler handles request from game instances. It responsibles of
// destroying game instance. It also responsible on sending output message
// to terminal/api by interprets commands from game instance.
func (s *server) outputHandler() {
	for {
		select {
		case gameOut := <-s.out:
			switch gameOut.command {
			case cmdGameDestroy:
				delete(s.games, gameOut.gameID)
			case cmdGamePrint:
				s.responder.print(gameOut.gameID)(&gameOut)
			}
		}
	}
}
