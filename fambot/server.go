package fambot

import (
	io_cli "github.com/pikomonde/fam100bot/io/client"
)

type handlerFunc func(*gameOutput)

// server contains info of all running games. This server only has 1
// instance and initialized 1 time.
type server struct {
	games map[string]*game
	out   chan gameOutput
	r     *responder
	cli   *io_cli.Client
}

type serverOpt struct {
	cli *io_cli.Client
}

// newServer initialize server, this should be only called once.
func newServer(opt serverOpt) *server {
	s := server{
		games: make(map[string]*game),
		out:   make(chan gameOutput),
		cli:   opt.cli,
		r: newResponder(responderOpt{
			cli: opt.cli,
		}),
	}
	go s.outputHandler()
	return &s
}

// inputHandler handles request from inputListener (which is api and/or
// webhook request). This instance universally interprets commands from
// any trigger.
func (s *server) inputHandler(uIn userInput) {
	gm, ok := s.games[uIn.gameID]
	switch uIn.command {
	case cmdUserJoin:
		// create new game if it doesn't exist
		if !ok {
			gm = s.newGame(uIn.gameID)
			gm.cli = s.cli
			s.games[uIn.gameID] = gm
			go gm.areaWaiting()
		}
		gm.in <- uIn
	case cmdUserAnswer:
		if ok {
			gm.in <- uIn
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
		case gOut := <-s.out:
			switch gOut.command {
			case cmdGameDestroy:
				delete(s.games, gOut.gameID)
			case cmdGamePrint:
				s.r.print(gOut.gameID)(&gOut)
			}
		}
	}
}
