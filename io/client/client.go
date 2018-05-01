package client

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pikomonde/fam100bot/io"
	"github.com/pikomonde/fam100bot/io/direct"
	"github.com/pikomonde/fam100bot/io/line"
)

// Client lists all clients used in the i/o
type Client struct {
	Direct *direct.Client
	Line   *line.Client
}

// New creates new universal client.
func New() *Client {
	return &Client{
		Line:   line.New(),
		Direct: direct.New(),
	}
}

// GetUserInput returns sets of information of UserInput.
func (cli *Client) GetUserInput(c *gin.Context, src string) io.UserInput {
	switch src {
	case io.SrcDir:
		return cli.Direct.GetUserInput(c)
	case io.SrcLine:
		return cli.Line.GetUserInput(c)
	}
	return io.UserInput{}
}

// GetFullName returns full name based on its userID.
func (cli *Client) GetFullName(userID string) (string, error) {
	uID := io.NewUserID(userID)
	if uID.Prefix != io.PreUser {
		errMsg := fmt.Sprintf("[GetFullName] "+
			"userID with prefix \"usr\" needed, "+
			"but ID with prefix \"%s\" found.\n", uID.Prefix)
		fmt.Print(errMsg)
		return "", errors.New(errMsg)
	}

	switch uID.Source {
	case io.SrcDir:
		return fmt.Sprintf("Player %s", uID.ID), nil
	case io.SrcLine:
		profile, err := cli.Line.Bot.GetProfile(uID.ID).Do()
		if err != nil {
			errMsg := fmt.Sprintf("[GetFullName][Line] "+
				"Error while getting profile from Line API. "+
				"Error: %s\n", err.Error())
			fmt.Print(errMsg)
			return "", errors.New(errMsg)
		}
		if len(profile.DisplayName) > 15 {
			profile.DisplayName = profile.DisplayName[:15]
		}
		return profile.DisplayName, nil
	}

	errMsg := fmt.Sprintf("[GetFullName] "+
		"Source unexpected, found \"%s\" instead.\n", uID.Source)
	fmt.Print(errMsg)
	return "", errors.New(errMsg)
}
