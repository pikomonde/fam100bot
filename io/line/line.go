package line

import (
	"fmt"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Module contains linebot
type Module struct {
	Client *linebot.Client
}

// New creates new linebot module.
func New() *Module {
	client := &http.Client{}
	bot, err := linebot.New(os.Getenv("chan_secret"),
		os.Getenv("chan_token"),
		linebot.WithHTTPClient(client))
	if err != nil {
		fmt.Println("[New] Can't create linebot client. Error:", err)
	}

	return &Module{
		Client: bot,
	}
}
