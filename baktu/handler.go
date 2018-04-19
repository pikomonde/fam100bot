package baktu

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	io_dir "github.com/pikomonde/fam100bot/io/direct"
)

// Module contains all config needed
type Module struct {
	router *gin.Engine
	server *server
	//DB *database.Store
}

// NewModule creates new input handler module.
func NewModule(router *gin.Engine) *Module {
	return &Module{
		router: router,
		server: newServer(),
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
	client := &http.Client{}
	bot, err := linebot.New(os.Getenv("chan_secret"), os.Getenv("chan_token"), linebot.WithHTTPClient(client))

	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		// Do something when something bad happened.
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			// Do Something...
			_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Piko Monde")).Do()
			if err != nil {
				// Do something when some bad happened
			}
		}
	}

	// ui := io_lne.GetUserInput(c)
	// m.server.inputHandler(userInput{
	// 	userID:  ui.UserID,
	// 	gameID:  ui.GameID,
	// 	command: ui.Command,
	// })

	// c.JSON(http.StatusOK, gin.H{
	// 	"status":  http.StatusOK,
	// 	"message": "direct",
	// })
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
