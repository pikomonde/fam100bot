package baktu

import (
	"fmt"
)

// server contains info of all running games. This server only has 1
// instance and initialized 1 time
type server struct {
	games map[string]*game
	out   chan gameOutput
}

// newServer initialize server, this should be only called once
func newServer() *server {
	var s server
	s.games = make(map[string]*game)
	s.out = make(chan gameOutput)
	go func() {
		for {
			select {
			case gameOut := <-s.out:
				switch gameOut.command {
				case cmdGameDestroy:
					delete(s.games, gameOut.gameID)
				case cmdGamePrint:
					fmt.Printf("%s> %s", gameOut.gameID, gameOut.message)
				}
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
		gm.in <- userIn
	case cmdUserHit:
		if ok {
			gm.in <- userIn
		}
	case cmdUserScore:
		gm.printScores()
	}
}
