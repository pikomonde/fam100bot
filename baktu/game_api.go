package baktu

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/pikomonde/fam100bot/io"
)

// join registers a player to a game, this function will return true if
// the player is successfully registered, return false if the player
// already registered before.
func (gm *game) join(userID string) bool {
	if _, ok := gm.players[userID]; !ok {
		fullname, err := gm.cli.GetFullName(userID)
		if err != nil {
			fullname = io.NewUserID(userID).ID
		}
		gm.players[userID] = &player{
			userID:    userID,
			gameScore: 0,
			fullname:  fullname,
		}
		return true
	}
	return false
}

// hit calculates scoring for a round, then store its value to roundScore.
func (gm *game) hit(userID string, startTime time.Time) {
	endMSec := int64(gm.mgAreaDur / time.Millisecond)
	nowMSec := int64(time.Now().Sub(startTime) / time.Millisecond)
	gm.players[userID].roundScore = endMSec - nowMSec
}

// isHitable returns boolean, which indicate whether a player can try
// another hit or not. It simply checking whether a player's roundScore
// is empty or not. An empty roundScore indicated the player never hit
// in a round before.
func (gm *game) isHitable(userID string) bool {
	return gm.players[userID].roundScore == 0
}

// giveGamePenalties gives penalties to the users that joining the game
// the middle of a game. It will add penalites which count how many
// rounds he pass.
func (gm *game) giveGamePenalties(userID string) {
	gm.players[userID].gameScore = int64(gm.mgNumRound-gm.mgRoundLeft) *
		int64(gm.mgAreaDur/time.Millisecond)
}

// giveRoundPenalties gives penalties to the users that hit over the time
// limit, or simply not hitting in a round.
func (gm *game) giveRoundPenalties() {
	for key := range gm.players {
		if gm.players[key].roundScore == 0 {
			gm.players[key].roundScore =
				int64(gm.mgAreaDur / time.Millisecond)
		}
	}
}

// addToTotalScore simply add roundScore to gameScore.
func (gm *game) addToTotalScore() {
	for key := range gm.players {
		gm.players[key].gameScore += gm.players[key].roundScore
	}
}

// resetRoundScore simply reset roundScore.
func (gm *game) resetRoundScore() {
	for key := range gm.players {
		gm.players[key].roundScore = 0
	}
}

// printScores prints player's roundScore and gameScore in sorted order.
func (gm *game) printScores(uIn ...userInput) {
	// sorting players by gameRound
	var ps players = make([]player, 0)
	for _, p := range gm.players {
		ps = append(ps, *p)
	}
	sort.Sort(ps)

	var msg string
	for _, v := range ps {
		msg += fmt.Sprintf("%s: %d (%d)\n",
			gm.players[v.userID].fullname,
			v.roundScore,
			v.gameScore)
	}
	gm.printf(strings.TrimSpace(msg))
}

// printf prints text to client (can be terminal, line, telegram, slack)
// by passing it to the "out" channel through gameOutput struct which
// also contains gameID of the message.
func (gm *game) printf(format string, a ...interface{}) {
	gm.out <- gameOutput{
		command: cmdGamePrint,
		gameID:  gm.gameID,
		message: fmt.Sprintf(format, a...),
	}
}

// rprintf behave same as printf, the only difference is, it also brings
// the userInput with it.
func (gm *game) rprintf(uIn userInput, format string, a ...interface{}) {
	gm.out <- gameOutput{
		command: cmdGamePrint,
		gameID:  gm.gameID,
		message: fmt.Sprintf(format, a...),
		uIn:     uIn,
	}
}
