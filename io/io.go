package io

// UserInput used by both cmdHandler and game to easily pass user input
// such as: GameID, UserID, and Commands.
type UserInput struct {
	GameID  string
	UserID  string
	Command int8
}

// GameOutput used by both server and game to easily pass game response
// such as: GameID, Message, and Commands.
type GameOutput struct {
	GameID  string
	Message string
	Command int8
}
