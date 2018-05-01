package line

import (
	"fmt"
	"net/http"
	"os"

	spew "github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/pikomonde/fam100bot/io"
)

// Client contains linebot client
type Client struct {
	Bot *linebot.Client
}

// New creates new linebot client.
func New() *Client {
	client := &http.Client{}
	cli, err := linebot.New(os.Getenv("chan_secret"),
		os.Getenv("chan_token"),
		linebot.WithHTTPClient(client))
	if err != nil {
		fmt.Println("[New] Can't create linebot client. Error:", err)
	}

	return &Client{
		Bot: cli,
	}
}

// GetUserInput returns sets of information of UserInput. It returns
// UserID, if it doesn't exist in the request it will return default
// UserID "8851". It also returns GameID, which will return default GameID
// "7411" if it doesn't exist. Lastly, it returns the Command code which
// will return 0 if whether the request not exist or in the wrong format.
func (c *Client) GetUserInput(ctx *gin.Context) (ui io.UserInput) {
	events, err := c.Bot.ParseRequest(ctx.Request)
	if err != nil {
		fmt.Println("[GetUserInput][ParseRequest] Can't parse request. "+
			"Err:", err)
	}

	var e *linebot.Event
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			e = event
		} else {
			// TODO: change dump to better location, such like, slack,
			// telegram, or even line.. It's better to use slack/telegram,
			// because line have api usage limitation 10k/month, meanwhile
			// telegram have 30 request/second limit, but slack (free
			// edition) have 10k message history.
			spew.Dump("[log][line][event]", event)
		}
	}

	if e != nil {
		ui.UserID = newUserID(e).String()
		ui.GameID = newGameID(e).String()
		ui.Command = getCommand(e)
	}
	return ui
}

// newGameID gets a correct format of gameID. It should be in format of
// "gme":[SRC]:[ID], where [SRC] is a const of request source and [ID] is
// string that contains number 0-9. This function returns io.GameID.
func newGameID(e *linebot.Event) *io.GameID {
	gID := io.GameID{}

	// part 1 should contains only the string "gme"
	gID.Prefix = io.PreGame

	// part 2 should contains source of the input
	gID.Source = io.SrcLine

	// part 3 should contains only string of number 0-9
	switch e.Source.Type {
	case "group":
		gID.ID = e.Source.GroupID
	case "room":
		gID.ID = e.Source.RoomID
	case "user":
		gID.ID = e.Source.UserID
	}

	return &gID
}

// newUserID gets a correct format of userID. It should be in "usr":[ID]
// format, where [ID] is string that contains only alphanumeric letter.
// This function returns io.UserID.
func newUserID(e *linebot.Event) *io.UserID {
	uID := io.UserID{}

	// part 1 should contains only the string "usr"
	uID.Prefix = io.PreUser

	// part 2 should contains source of the input
	uID.Source = io.SrcLine

	// part 3 should contains only alphanumeric string
	uID.ID = e.Source.UserID

	return &uID
}

// getCommand gets command code from context.
func getCommand(e *linebot.Event) string {
	switch msg := e.Message.(type) {
	case *linebot.TextMessage:
		// cmd, err := strconv.ParseInt(msg.Text, 10, 64)
		// if err != nil {
		// 	return 0
		// }
		// return int8(cmd)
		return msg.Text
	}
	return ""
}
