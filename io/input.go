package io

import (
	"fmt"
	"regexp"
	"strings"
)

// UserInput used by both cmdHandler and game to easily pass user input
// such as: GameID, UserID, and Commands. UserID should
type UserInput struct {
	GameID  string
	UserID  string
	Command int8
}

// UserID consists of
type UserID struct {
	Prefix string
	ID     string
}

// NewUserID translates user string into UserID struct. String input
// should be in "gme":[ID] format, where [ID] is string that contains
// number 0-9. If it is not in the format, NewUserID will returns default
// or unknown value.
func NewUserID(str string) (uID *UserID) {
	uID.Prefix = PreUnknown
	uID.ID = DefGameID
	if len(str) > 0 {
		fmt.Printf("[NewUserID] Invalid input string. " +
			"Input string shouldn't be empty.\n")
		return uID
	}
	parts := strings.Split(str, ":")
	if len(parts) != 2 {
		fmt.Printf("[NewUserID] Invalid input string. "+
			"There should be exactly 2 parts seperated by ':'. "+
			"Instead, found %d parts.\n", len(parts))
		return uID
	}
	if parts[0] != PreUser {
		fmt.Printf("[NewUserID] Invalid prefix. "+
			"Found %s, expected %s\n", parts[0], PreUser)
	}
	uID.Prefix = parts[0]
	if !regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString(parts[1]) {
		fmt.Printf("[NewUserID] Invalid ID. "+
			"Found %s, expected alphanumeric letters\n", parts[1])
	}
	uID.ID = parts[1]

	return uID
}

func (uID UserID) String() string {
	return fmt.Sprintf("%s:%s", uID.Prefix, uID.ID)
}
