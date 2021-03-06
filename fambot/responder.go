package fambot

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"

	"github.com/pikomonde/fam100bot/io"
	io_cli "github.com/pikomonde/fam100bot/io/client"
)

// responder contains all config needed
type responder struct {
	cli *io_cli.Client
}

type responderOpt struct {
	cli *io_cli.Client
}

// newResponder creates new output handler responder. This responder only
// has 1 instance and initialized 1 time.
func newResponder(opt responderOpt) *responder {
	return &responder{
		cli: opt.cli,
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
	fmt.Printf("Room %s>>\n%s\n",
		io.NewGameID(gOut.gameID).ID,
		gOut.message)
}

func (r *responder) linebot(gOut *gameOutput) {
	r.cli.Line.Bot.PushMessage(
		io.NewGameID(gOut.gameID).ID,
		linebot.NewTextMessage(gOut.message),
	).Do()
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
