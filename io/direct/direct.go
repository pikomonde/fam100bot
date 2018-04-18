package direct

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pikomonde/fam100bot/io"
)

func GetUserInput(c *gin.Context) (ui io.UserInput) {
	ui.UserID = c.DefaultQuery("user_id", "8851")
	ui.GameID = io.SrcDir + c.DefaultQuery("game_id", "7411")
	cmd, _ := strconv.ParseInt(c.DefaultQuery("command", "0"), 10, 8)
	ui.Command = int8(cmd)
	return ui
}
