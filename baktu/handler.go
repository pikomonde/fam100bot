package baktu

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
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
	ui := io_dir.GetUserInput(c)
	m.server.inputHandler(userInput{
		userID:  ui.UserID,
		gameID:  ui.GameID,
		command: ui.Command,
	})

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "direct",
	})
}

func (m *Module) lineWebhook(c *gin.Context) {
	cli := m.server.responder.line.Client
	events, err := cli.ParseRequest(c.Request)
	if err != nil {
		// Do something when something bad happened.
	}
	fmt.Println("==== TEST 1")
	var gameID = "ABC123xyz"
	var cmd int64 = 4
	for _, event := range events {
		fmt.Println("==== TEST 2")
		if event.Type == linebot.EventTypeMessage {
			fmt.Println("==== TEST 3", event.Source.Type)
			switch event.Source.Type {
			case "group":
				gameID = event.Source.GroupID
				fmt.Println(event.Source.Type, gameID)
			case "room":
				gameID = event.Source.RoomID
				fmt.Println(event.Source.Type, gameID)
			case "user":
				gameID = event.Source.UserID
				fmt.Println(event.Source.Type, gameID)
			}
			switch msg := event.Message.(type) {
			case *linebot.TextMessage:
				cmd, _ = strconv.ParseInt(msg.Text, 10, 64)
			}
		}
	}

	ui := io_dir.GetUserInput(c)
	m.server.inputHandler(userInput{
		userID:  ui.UserID,
		gameID:  "gme:lne:" + gameID,
		command: int8(cmd),
	})

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "direct",
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
