package baktu

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	io_dir "github.com/pikomonde/fam100bot/io/direct"
	io_lne "github.com/pikomonde/fam100bot/io/line"
)

// Option contains all config needed
type Option struct {
	LineModule *io_lne.Module
}

// Module contains all config needed
type Module struct {
	router *gin.Engine
	server *server
	//DB *database.Store
}

// New creates new input handler module.
func New(router *gin.Engine, opt Option) *Module {
	return &Module{
		router: router,
		server: newServer(serverOpt{
			line: opt.LineModule,
		}),
	}
}

// Register the endpoints.
func (m *Module) Register() {
	m.router.GET("/baktu/ping", m.ping)
	m.router.POST("/baktu/direct", m.direct)
	m.router.POST("/baktu/line/webhook", m.lineWebhook)
}

func (m *Module) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "baktu pong",
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

	ui := io_dir.GetUserInput(c)
	m.server.inputHandler(userInput{
		userID:  ui.UserID,
		gameID:  ui.GameID,
		command: ui.Command,
	})
	fmt.Println("[log][direct]", ui)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "direct",
	})
}

func (m *Module) lineWebhook(c *gin.Context) {
	cli := m.server.responder.line.Client
	ui := io_lne.GetUserInput(c, cli)
	m.server.inputHandler(userInput{
		userID:  ui.UserID,
		gameID:  ui.GameID,
		command: ui.Command,
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
}

// consts of cmdUser is a command signal that used by userInput.
const (
	// cmdUserJoin used whenever a player want to join a game.
	cmdUserJoin = int8(iota)
	// cmdUserHit used whenever a player want to guess the time in a
	// gameplay.
	cmdUserHit
	// cmdUserScore used whenever a player want to check their overall
	// total score.
	cmdUserScore
)
