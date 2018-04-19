package baktu

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"

	"github.com/pikomonde/fam100bot/io"
	io_lne "github.com/pikomonde/fam100bot/io/line"
)

// responder contains all config needed
type responder struct {
	line *io_lne.Module
}

type responderOpt struct {
	line *io_lne.Module
}

// newResponder creates new output handler responder. This responder only
// has 1 instance and initialized 1 time.
func newResponder(opt responderOpt) *responder {
	return &responder{
		line: opt.line,
	}
}

func (r *responder) print(gID string) handlerFunc {
	switch io.NewGameID(gID).Source {
	case io.SrcLine:
		return r.linebot
	case io.SrcDir:
		return r.terminal
	case io.SrcUnknown:
		return func(*gameOutput) {}
	}
	return func(*gameOutput) {}
}

func (r *responder) terminal(gOut *gameOutput) {
	fmt.Printf("Room %s> %s",
		io.NewGameID(gOut.gameID).ID,
		gOut.message)
}

func (r *responder) linebot(gOut *gameOutput) {
	msg := fmt.Sprintf("Room %s> %s",
		io.NewGameID(gOut.gameID).ID,
		gOut.message)

	cli := r.line.Client
	cli.PushMessage(io.NewGameID(gOut.gameID).ID, linebot.NewTextMessage(msg)).Do()
	fmt.Println(msg)
}

// gameOutput used by both server and game to easily pass game response
// such as: gameID, message, and commands.
type gameOutput struct {
	uIn     userInput
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
