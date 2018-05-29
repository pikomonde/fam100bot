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
	Command string
}

// UserID consists of a prefix "usr" and [ID] is string that contains
// number 0-9.
type UserID struct {
	Prefix string
	Source string
	ID     string
}

// NewUserID translates user string into UserID struct. String input
// should be in "usr":[SRC]:[ID] format, where [SRC] is a const of request
// source and [ID] is string that contains number 0-9. If it is not in the
// format, NewUserID will returns default or unknown value.
func NewUserID(str string) *UserID {
	uID := UserID{
		Prefix: PreUnknown,
		Source: SrcUnknown,
		ID:     DefGameID,
	}
	if len(str) == 0 {
		fmt.Printf("[NewUserID] Invalid input string. " +
			"Input string shouldn't be empty.\n")
		return &uID
	}
	parts := strings.Split(str, ":")
	if len(parts) != 3 {
		fmt.Printf("[NewUserID] Invalid input string. "+
			"There should be exactly 3 parts seperated by ':'. "+
			"Instead, found %d parts in %s\n", len(parts), str)
		return &uID
	}
	if parts[0] != PreUser {
		fmt.Printf("[NewUserID] Invalid prefix. "+
			"Found %s, expected %s\n", parts[0], PreUser)
	}
	uID.Prefix = parts[0]
	if !regexp.MustCompile(`^[a-z]{3}$`).MatchString(parts[1]) {
		fmt.Printf("[NewUserID] Invalid source. "+
			"Found %s, expected [ABC]\n", parts[1])
	}
	uID.Source = parts[1]
	if !regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString(parts[2]) {
		fmt.Printf("[NewUserID] Invalid ID. "+
			"Found %s, expected alphanumeric letters\n", parts[2])
	}
	uID.ID = parts[2]

	return &uID
}

// String stringify userID from object type into string.
func (uID UserID) String() string {
	return fmt.Sprintf("%s:%s:%s", uID.Prefix, uID.Source, uID.ID)
}
