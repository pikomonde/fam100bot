package baktu

import (
	"fmt"
	"time"
)

// join registers a player to a game, this function will return true if
// the player is successfully registered, return false if the player
// already registered before.
func (gm *game) join(userID string) bool {
	if _, ok := gm.players[userID]; !ok {
		gm.players[userID] = &player{
			userID:    userID,
			gameScore: 0,
		}
		return true
	}
	return false
}

// hit calculates scoring for a round, then store its value to roundScore
func (gm *game) hit(userID string, startTime time.Time) {
	endMSec := int64(gm.mgAreaDur / time.Millisecond)
	nowMSec := int64(time.Now().Sub(startTime) / time.Millisecond)
	gm.players[userID].roundScore = endMSec - nowMSec
}

// isHitable returns boolean, which indicate whether a player can try
// another hit or not. It simply checking whether a player's roundScore
// is empty or not. An empty roundScore indicated the player never hit
// in a round before
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

func (gm *game) addToTotalScore() {
	for key := range gm.players {
		gm.players[key].gameScore += gm.players[key].roundScore
	}
}

func (gm *game) resetRoundScore() {
	for key := range gm.players {
		gm.players[key].roundScore = 0
	}
}

func (gm *game) printScores() {
	// TODO: sort it first
	for _, v := range gm.players {
		fmt.Printf("Player %s: %d (%d)\n",
			v.userID,
			v.roundScore,
			v.gameScore)
	}
}
