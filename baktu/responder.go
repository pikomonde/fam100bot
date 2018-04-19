package baktu

import (
	"fmt"

	"github.com/pikomonde/fam100bot/io"
)

// responder contains all config needed
type responder struct{}

// newResponder creates new output handler responder.
func newResponder() *responder {
	return &responder{}
}

func (r *responder) print(gID string) handlerFunc {
	switch io.NewGameID(gID).Source {
	case io.SrcLine:
		return func(*gameOutput) {}
	case io.SrcDir:
		return r.terminal
	case io.SrcUnknown:
		return func(*gameOutput) {}
	}
	return func(*gameOutput) {}
}

func (r *responder) terminal(gameOut *gameOutput) {
	fmt.Printf("%s> %s", gameOut.gameID, gameOut.message)
}

// gameOutput used by both server and game to easily pass game response
// such as: gameID, message, and commands.
type gameOutput struct {
	gameID  string
	message string
	command int8
}

// consts of cmdGame is a command signal that used by gameOutput
const (
	// cmdGameDestroy used whenever a game want to destroy itself. Usually
	// a game want to destroy itself when the game ended or a game done
	// waiting for players.
	cmdGameDestroy = int8(iota)
	// cmdGamePrint used whenever a game want to print the output message
	// to a medium (terminal, line, telegram, or slack).
	cmdGamePrint
)
