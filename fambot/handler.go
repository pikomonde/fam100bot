package fambot

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pikomonde/fam100bot/io"
	io_cli "github.com/pikomonde/fam100bot/io/client"
)

// Option contains all config needed
type Option struct {
	Client *io_cli.Client
}

// Module contains all config needed
type Module struct {
	router *gin.Engine
	s      *server
	cli    *io_cli.Client
	//DB *database.Store
}

// New creates new input handler module.
func New(router *gin.Engine, opt Option) *Module {
	return &Module{
		router: router,
		cli:    opt.Client,
		s: newServer(serverOpt{
			cli: opt.Client,
		}),
	}
}

// Register the endpoints.
func (m *Module) Register() {
	m.router.GET("/fambot/ping", m.ping)
	m.router.POST("/fambot/direct", m.direct)
	m.router.POST("/fambot/line/webhook", m.lineWebhook)
}

func (m *Module) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "fambot pong",
	})
}

func (m *Module) direct(c *gin.Context) {
	if os.Getenv("env") == "prod" {
		c.JSON(http.StatusForbidden, gin.H{
			"status":  http.StatusForbidden,
			"message": "Forbidden",
		})
		return
	}

	ui := m.cli.GetUserInput(c, io.SrcDir)

	var cmd = cmdUserJoin
	var msg string
	switch strings.ToLower(ui.Command) {
	case "0", "join":
		cmd = cmdUserJoin
	case "2", "score":
		cmd = cmdUserScore
	default:
		cmd = cmdUserAnswer
		msg = strings.ToLower(ui.Command)
	}
	m.s.inputHandler(userInput{
		userID:  ui.UserID,
		gameID:  ui.GameID,
		command: cmd,
		message: msg,
	})
	fmt.Println("[log][direct]", ui)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "direct",
	})
}

func (m *Module) lineWebhook(c *gin.Context) {
	ui := m.cli.GetUserInput(c, io.SrcLine)

	var cmd = cmdUserJoin
	var msg string
	switch strings.ToLower(ui.Command) {
	case "join":
		cmd = cmdUserJoin
	case "score":
		cmd = cmdUserScore
	default:
		cmd = cmdUserAnswer
		msg = strings.ToLower(ui.Command)
	}
	m.s.inputHandler(userInput{
		userID:  ui.UserID,
		gameID:  ui.GameID,
		command: cmd,
		message: msg,
	})
	fmt.Println("[log][line]", ui)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "line",
	})
}

// userInput used by both cmdHandler and game to easily pass user input
// such as: gameID, userID, and commands.
type userInput struct {
	gameID  string
	userID  string
	command int8
	message string
}

// consts of cmdUser is a command signal that used by userInput.
const (
	// cmdUserJoin used whenever a player want to join a game.
	cmdUserJoin = int8(iota)
	// cmdUserHit used whenever a player want to guess the time in a
	// gameplay.
	cmdUserAnswer
	// cmdUserScore used whenever a player want to check their overall
	// total score.
	cmdUserScore
)
