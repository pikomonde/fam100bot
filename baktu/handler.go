package baktu

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (m *Module) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "baktu pong",
	})
}

func (m *Module) direct(c *gin.Context) {
	// TODO: get this gameID from somewhere (roomID, anything)
	userID := c.DefaultQuery("user_id", "usr8851")
	gameID := c.DefaultQuery("game_id", "gm7411")
	command, _ := strconv.ParseInt(c.DefaultQuery("command", "0"), 10, 8)
	m.server.inputHandler(userInput{
		userID:  userID,
		gameID:  gameID,
		command: int8(command),
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

// gameOutput used by both server and game to easily pass game response
// such as: gameID, message, and commands.
type gameOutput struct {
	gameID  string
	message string
	command int8
}
