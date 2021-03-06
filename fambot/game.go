package fambot

import (
	"os"
	"time"

	io_cli "github.com/pikomonde/fam100bot/io/client"
	"github.com/pikomonde/fam100bot/pfmt"
)

// This "fambot" game consists of 3 parts:
// 1. The waiting area
// 2. The main game area
// 3. The break area
// Each of these partsare a goroutine, which means after created, it can
// be called only from "in" channel. Please make sure to close (return)
// this concurrency by using timeout. To destroy the entire game instance
// (all these 3 parts + all game data), it can be done easily by using
// commanding destroy to "out" channel.

// game contains all variables a game needed. This game "object" only
// created when a player joining a game. It will be destroyed whenever all
// player out from the game or the game ends.
type game struct {
	gameID    string // The game ID
	in        chan userInput
	out       chan gameOutput
	players   map[string]*player
	questions []*question
	cli       *io_cli.Client

	wNotifyDur  time.Duration
	wAreaDur    time.Duration
	wMinPlayer  int8
	bAreaDur    time.Duration
	mgRoundLeft int8
	mgNumRound  int8 // should be a const, not changed
	mgNotifyDur time.Duration
	mgAreaDur   time.Duration
}

// newGame created a new game (in a game room). There can only be maximum
// 1 game instance running in a room.
func (s *server) newGame(gameID string) *game {
	if os.Getenv("env") == "dev" {
		return &game{
			gameID:      gameID,
			in:          make(chan userInput),
			out:         s.out,
			players:     make(map[string]*player),
			questions:   s.newQuestions(),
			wNotifyDur:  10 * time.Second,
			wAreaDur:    120 * time.Second,
			wMinPlayer:  2,
			bAreaDur:    5 * time.Second,
			mgRoundLeft: 5,
			mgNumRound:  5,
			mgNotifyDur: 10 * time.Second,
			mgAreaDur:   90 * time.Second,
		}
	}
	return &game{
		gameID:      gameID,
		in:          make(chan userInput),
		out:         s.out,
		players:     make(map[string]*player),
		questions:   s.newQuestions(),
		wNotifyDur:  10 * time.Second,
		wAreaDur:    120 * time.Second,
		wMinPlayer:  3,
		bAreaDur:    5 * time.Second,
		mgRoundLeft: 3,
		mgNumRound:  3,
		mgNotifyDur: 10 * time.Second,
		mgAreaDur:   90 * time.Second,
	}
}

// areaWaiting is used for waiting player to join the game. If the minimum
// requirement player is not met under certain area time, the game will be
// closed (and the game instance will be destroyed). A player needs to
// rejoin the game in order to play.
func (gm *game) areaWaiting() {
	startTime := time.Now()
	timeout := time.After(gm.wAreaDur)
	notify := time.Tick(gm.wNotifyDur)
	for {
		select {
		case uIn := <-gm.in:
			if uIn.command == cmdUserJoin {
				ok := gm.join(uIn.userID)
				if int8(len(gm.players)) >= gm.wMinPlayer {
					gm.rprintf(uIn,
						"%s bergabung. Permainan dimulai..",
						gm.players[uIn.userID].fullname)
					go gm.areaBreak()
					return
				}
				if ok {
					endSec := int(gm.wAreaDur / time.Second)
					nowSec := int(time.Now().Sub(startTime) / time.Second)
					gm.rprintf(uIn,
						"%s bergabung. Menunggu %d orang lagi. "+
							"Menunggu %s detik sebelum berakhir!",
						gm.players[uIn.userID].fullname,
						gm.wMinPlayer-int8(len(gm.players)),
						pfmt.Time(endSec-nowSec))
				}
			}
		case <-timeout:
			gm.printf("Permainan dibatalkan, jumlah pemain kurang.")
			gm.out <- gameOutput{
				command: cmdGameDestroy,
				gameID:  gm.gameID,
			}
			return
		case <-notify:
			endSec := int(gm.wAreaDur / time.Second)
			nowSec := int(time.Now().Sub(startTime) / time.Second)
			switch endSec - nowSec {
			case 60, 30:
				gm.printf("%d detik lagi!", endSec-nowSec)
			}
		}
	}
}

// areaMainGame is an area where the main game is occured. It does
// countdown from a certain time. It excpect a player to "ANSWER" before
// the countdown reach 0.
func (gm *game) areaMainGame() {
	startTime := time.Now()
	timeout := time.After(gm.mgAreaDur)
	notify := time.Tick(10 * time.Second)
	for {
		select {
		case uIn := <-gm.in:
			// player automatically join the game, whether they command
			// "JOIN" or "ANSWER"
			gm.join(uIn.userID)
			if uIn.command == cmdUserAnswer {
				if gm.answer(uIn.userID, uIn.message) {
					if gm.isAllAnswered() {
						gm.mainGameEnds()
						return
					}
					gm.printAnswers()
				}
			}
		case <-timeout:
			gm.mainGameEnds()
			return
		case <-notify:
			endSec := int(gm.mgAreaDur / time.Second)
			nowSec := int(time.Now().Sub(startTime) / time.Second)
			switch endSec - nowSec {
			case 60, 30:
				gm.printf("%d detik lagi!", endSec-nowSec)
			}
		}
	}
}

// mainGameEnds parts of areaMainGame for the End Game case
func (gm *game) mainGameEnds() {
	gm.addToTotalScore()
	gm.printf("Ronde %d berakhir", gm.round()+1)
	gm.printAllAnswers()
	gm.printScores()
	gm.resetRoundScore()
	gm.mgRoundLeft--
	if gm.mgRoundLeft > 0 {
		go gm.areaBreak()
	} else {
		gm.printf("Permainan berakhir...")
		gm.out <- gameOutput{
			command: cmdGameDestroy,
			gameID:  gm.gameID,
		}
	}
}

// areaBreak is used to give a break to players from each round. It also
// allows a new player to join the game, whenever a player command "JOIN"
// or "ANSWER".
func (gm *game) areaBreak() {
	timeout := time.After(gm.bAreaDur)
	gm.printf("Game akan segea di mulai!\n" +
		"Semua member chat boleh langsung menjawab tanpa \"join\".")
	for {
		select {
		case uIn := <-gm.in:
			// player automatically join the game, whether they command
			// "JOIN" or "HIT"
			gm.join(uIn.userID)
		case <-timeout:
			gm.printQuestions()
			go gm.areaMainGame()
			return
		}
	}
}
