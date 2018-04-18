package baktu

import (
	"time"
)

// This "baktu" game consists of 3 parts:
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
	gameID  string // The game ID
	in      chan userInput
	out     chan gameOutput
	players map[string]*player

	wNotifyDur  time.Duration
	wAreaDur    time.Duration
	wMinPlayer  int8
	bAreaDur    time.Duration
	mgRoundLeft int8
	mgNumRound  int8 // should be a const, not changed
	mgAreaDur   time.Duration
}

// newGame created a new game (in a game room). There can only be maximum
// 1 game instance running in a room.
func (s *server) newGame(gameID string) *game {
	return &game{
		gameID:      gameID,
		in:          make(chan userInput),
		out:         s.out,
		players:     make(map[string]*player),
		wNotifyDur:  10 * time.Second,
		wAreaDur:    30 * time.Second,
		wMinPlayer:  3,
		bAreaDur:    5 * time.Second,
		mgRoundLeft: 5,
		mgNumRound:  5,
		mgAreaDur:   10 * time.Second,
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
		case userIn := <-gm.in:
			if userIn.command == cmdUserJoin {
				ok := gm.join(userIn.userID)
				if int8(len(gm.players)) >= gm.wMinPlayer {
					gm.printf("%s bergabung. Permainan dimulai..\n",
						userIn.userID)
					go gm.areaBreak()
					return
				}
				if ok {
					gm.printf("%s bergabung. Menunggu %d orang lagi\n",
						userIn.userID,
						gm.wMinPlayer-int8(len(gm.players)))
				}
			}
		case <-timeout:
			gm.printf("Permainan dibatalkan, jumlah pemain kurang\n")
			gm.out <- gameOutput{
				command: cmdGameDestroy,
				gameID:  gm.gameID,
			}
			return
		case <-notify:
			endSec := int(gm.wAreaDur / time.Second)
			nowSec := int(time.Now().Sub(startTime) / time.Second)
			gm.printf("%d detik lagi!\n", endSec-nowSec)
		}
	}
}

// areaMainGame is an area where the main game is occured. It does
// countdown from a certain time. It excpect a player to "HIT" just before
// the countdown reach 0. The nearer the player "HIT" before it ends, the
// lower the score will be. Player will get score penalty whenver they
// "HIT" after the countdown finished. The lower the score, the better
// the player.
func (gm *game) areaMainGame() {
	startTime := time.Now()
	timeout := time.After(gm.mgAreaDur)
	notify := time.Tick(1 * time.Second)
	for {
		select {
		case userIn := <-gm.in:
			// player automatically join the game, whether they command
			// "JOIN" or "HIT"
			if gm.join(userIn.userID) {
				gm.giveGamePenalties(userIn.userID)
			}
			if userIn.command == cmdUserHit &&
				gm.isHitable(userIn.userID) {
				gm.hit(userIn.userID, startTime)
			}
		case <-timeout:
			gm.giveRoundPenalties()
			gm.addToTotalScore()
			gm.mgRoundLeft--
			gm.printf("Ronde %d berakhir\n",
				gm.mgNumRound-gm.mgRoundLeft)
			gm.printScores()
			gm.resetRoundScore()
			if gm.mgRoundLeft > 0 {
				go gm.areaBreak()
			} else {
				gm.printf("Permainan berakhir...\n")
				gm.out <- gameOutput{
					command: cmdGameDestroy,
					gameID:  gm.gameID,
				}
			}
			return
		case <-notify:
			// TODO: hide printing to create challenging game
			endSec := int(gm.mgAreaDur / time.Second)
			nowSec := int(time.Now().Sub(startTime) / time.Second)
			gm.printf("%d.\n", endSec-nowSec)
		}
	}
}

// areaBreak is used to give a break to players from each round. It also
// allows a new player to join the game, whenever a player command "JOIN"
// or "HIT".
func (gm *game) areaBreak() {
	timeout := time.After(gm.bAreaDur)
	gm.printf("Chat ketika angka menyentuh 0!\n")
	for {
		select {
		case userIn := <-gm.in:
			// player automatically join the game, whether they command
			// "JOIN" or "HIT"
			gm.join(userIn.userID)
		case <-timeout:
			go gm.areaMainGame()
			return
		}
	}
}
