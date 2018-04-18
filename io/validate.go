package io

import (
	"errors"
	"regexp"
	"strings"
)

// UserID verify wheteher it is a correct format of userID or not. userID
// should be in "usr":[ID] format, where [ID] is string that contains
// number 0-9.
func UserID(uID string) error {
	if uID == "" {
		return errors.New("userID should not be empty")
	}

	parts := strings.Split(uID, ":")
	if len(parts) != 2 {
		return errors.New("There should be 2 parts for userID")
	}

	if parts[0] != "usr" {
		return errors.New("1st part should contain 'usr'")
	}

	var validID = regexp.MustCompile(`^[0-9]+$`)
	if !validID.MatchString(parts[1]) {
		return errors.New("2nd part should only contain number 0-9")
	}

	return nil
}

// GameID verify wheteher it is a correct format of gameID or not. gameID
// should be in "gme":[SRC]:[ID] format, where [SRC] is a const of request
// source and [ID] is string that contains number 0-9.
func GameID(gID string) error {
	if gID == "" {
		return errors.New("gameID should not be empty")
	}

	parts := strings.Split(gID, ":")
	if len(parts) != 3 {
		return errors.New("There should be 3 parts for gameID")
	}

	if parts[0] != "gme" {
		return errors.New("1st part should contain 'gme'")
	}

	var validID = regexp.MustCompile(`^[A-Z]{3}$`)
	if !validID.MatchString(parts[1]) {
		return errors.New("2nd part should only contain 3 letters of A-Z")
	}

	validID = regexp.MustCompile(`^[0-9]+$`)
	if !validID.MatchString(parts[2]) {
		return errors.New("3rd part should only contain number 0-9")
	}

	return nil
}

const (
	// SrcUnknown indicates source are unknown
	SrcUnknown = "nil"
	// SrcDir indicates source are from direct api call
	SrcDir = "dir"
	// SrcLine indicates source are from line webhook
	SrcLine = "lne"
)

// GetSourceByGameID getting sourceID by parsing from GameID
func GetSourceByGameID(gameID string) string {
	if len(gameID) > 3 {
		switch gameID[0:3] {
		case SrcDir:
			return SrcDir
		case SrcLine:
			return SrcLine
		}
	}
	return SrcUnknown
}
