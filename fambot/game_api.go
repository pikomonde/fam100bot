package fambot

import (
	"fmt"
	"sort"

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
func (gm *game) answer(userID string, answer string) bool {
	round := int(gm.round())
	if round < 0 && round > len(gm.questions) {
		return false
	}

	a := gm.questions[round].answer(userID, answer)
	if a != nil {
		gm.players[userID].roundScore += a.Score
		return true
	}

	return false
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
	gm.printf(msg)
}

// printAllAnswers prints player's answer in a round with it's score.
func (gm *game) printAllAnswers(uIn ...userInput) {
	var q = gm.questions[gm.round()]
	var msg = fmt.Sprintf("[ID: %d] %s\n\n", q.QuestionID, q.Text)

	for i, a := range q.Answers {
		msg += fmt.Sprintf("%d. (%d) %s\n", i+1, a.Score, a.Text)
		if a.UserID != "" {
			msg += fmt.Sprintf("   %s\n", gm.players[a.UserID].fullname)
		}
	}
	gm.printf(msg)
}

// printAnswers prints player's answer in a round with it's score.
func (gm *game) printAnswers(uIn ...userInput) {
	var q = gm.questions[gm.round()]
	var msg = fmt.Sprintf("[ID: %d] %s\n\n", q.QuestionID, q.Text)

	for i, a := range q.Answers {
		if a.UserID != "" {
			msg += fmt.Sprintf("%d. (%d) %s\n", i+1, a.Score, a.Text)
			msg += fmt.Sprintf("   %s\n", gm.players[a.UserID].fullname)
		} else {
			msg += fmt.Sprintf("%d. ____________________\n", i+1)
		}
	}
	gm.printf(msg)
}

// printQuestions prints game question in a round with player's score.
func (gm *game) printQuestions(uIn ...userInput) {
	var q = gm.questions[gm.round()]
	var msg = fmt.Sprintf(
		"Ronde %d dari %d.\n\n"+
			"[ID: %d] %s\n\n",
		gm.round()+1, gm.mgNumRound, q.QuestionID, q.Text)

	for i := 1; i <= len(q.Answers); i++ {
		msg += fmt.Sprintf("%d. ____________________\n", i)
	}
	gm.printf(msg)
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

// round getting round number. It's values started from 0 (0, 1, 2, 3...).
func (gm *game) round() int8 {
	return gm.mgNumRound - gm.mgRoundLeft
}

// isAllAnswered checks whether all answer has been guessed or not.
func (gm *game) isAllAnswered() bool {
	var q = gm.questions[gm.round()]
	for _, a := range q.Answers {
		if a.UserID == "" {
			return false
		}
	}
	return true
}
