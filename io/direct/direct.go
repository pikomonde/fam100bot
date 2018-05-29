package direct

import (
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/pikomonde/fam100bot/io"
)

// Client contains "direct" i/o information
type Client struct{}

// New creates new terminal client.
func New() *Client {
	return &Client{}
}

// GetUserInput returns sets of information of UserInput. It returns
// UserID, if it doesn't exist in the request it will return default
// UserID "8851". It also returns GameID, which will return default GameID
// "7411" if it doesn't exist. Lastly, it returns the Command code which
// will return 0 if whether the request not exist or in the wrong format.
func (c *Client) GetUserInput(ctx *gin.Context) (ui io.UserInput) {
	ui.UserID = newUserID(ctx).String()
	ui.GameID = newGameID(ctx).String()
	ui.Command = getCommand(ctx)
	return ui
}

// newGameID gets a correct format of gameID. It should be in format of
// "gme":[SRC]:[ID], where [SRC] is a const of request source and [ID] is
// string that contains number 0-9. This function returns io.GameID.
func newGameID(ctx *gin.Context) *io.GameID {
	gID := io.GameID{}

	// part 1 should contains only the string "gme"
	gID.Prefix = io.PreGame

	// part 2 should contains source of the input
	gID.Source = ctx.DefaultQuery("source", io.SrcDir)
	if !regexp.MustCompile(`^[a-z]{3}$`).MatchString(gID.ID) {
		gID.ID = io.SrcUnknown
	}

	// part 3 should contains only string of number 0-9
	gID.ID = ctx.DefaultQuery("game_id", io.DefGameID)
	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(gID.ID) {
		gID.ID = io.DefGameID
	}

	return &gID
}

// newUserID gets a correct format of userID. It should be in format of
// "usr":[SRC]:[ID], where [SRC] is a const of request source and [ID] is
// string that contains number 0-9. This function returns io.UserID.
func newUserID(ctx *gin.Context) *io.UserID {
	uID := io.UserID{}

	// part 1 should contains only the string "usr"
	uID.Prefix = io.PreUser

	// part 2 should contains source of the input
	uID.Source = ctx.DefaultQuery("source", io.SrcDir)
	if !regexp.MustCompile(`^[a-z]{3}$`).MatchString(uID.ID) {
		uID.ID = io.SrcUnknown
	}

	// part 3 should contains only string of number 0-9
	uID.ID = ctx.DefaultQuery("user_id", io.DefUserID)
	if !regexp.MustCompile(`^[0-9]+$`).MatchString(uID.ID) {
		uID.ID = io.DefUserID
	}

	return &uID
}

// getCommand gets command code from context.
func getCommand(ctx *gin.Context) string {
	// cmd, err := strconv.ParseInt(ctx.DefaultQuery("command", "0"), 10, 8)
	// if err != nil {
	// 	return 0
	// }
	// return int8(cmd)
	return ctx.DefaultQuery("command", "")
}
