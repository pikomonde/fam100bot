package io

import (
	"fmt"
	"regexp"
	"strings"
)

// GameOutput used by both server and game to easily pass game response
// such as: GameID, Message, and Commands.
type GameOutput struct {
	GameID  string
	Message string
	Command int8
}

// GameID consists of a prefix "gme", source which is a const of request
// source and [ID] is string that contains number 0-9.
type GameID struct {
	Prefix string
	Source string
	ID     string
}

// NewGameID translates game string into GameID struct. String input
// should be in "gme":[SRC]:[ID] format, where [SRC] is a const of request
// source and [ID] is string that contains number 0-9. If it is not in the
// format, NewGameID will returns default or unknown value.
func NewGameID(str string) *GameID {
	gID := GameID{
		Prefix: PreUnknown,
		Source: SrcUnknown,
		ID:     DefGameID,
	}
	if len(str) == 0 {
		fmt.Printf("[NewGameID] Invalid input string. " +
			"Input string shouldn't be empty.\n")
		return &gID
	}
	parts := strings.Split(str, ":")
	if len(parts) != 3 {
		fmt.Printf("[NewGameID] Invalid input string. "+
			"There should be exactly 3 parts seperated by ':'. "+
			"Instead, found %d parts in %s\n", len(parts), str)
		return &gID
	}
	if parts[0] != PreGame {
		fmt.Printf("[NewGameID] Invalid prefix. "+
			"Found %s, expected %s\n", parts[0], PreGame)
	}
	gID.Prefix = parts[0]
	if !regexp.MustCompile(`^[a-z]{3}$`).MatchString(parts[1]) {
		fmt.Printf("[NewGameID] Invalid source. "+
			"Found %s, expected [ABC]\n", parts[1])
	}
	gID.Source = parts[1]
	if !regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString(parts[2]) {
		fmt.Printf("[NewGameID] Invalid ID. "+
			"Found %s, expected alphanumeric letters\n", parts[2])
	}
	gID.ID = parts[2]

	return &gID
}

func (gID *GameID) String() string {
	return fmt.Sprintf("%s:%s:%s", gID.Prefix, gID.Source, gID.ID)
}
