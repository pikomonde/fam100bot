package baktu

type player struct {
	userID     string
	gameScore  int64
	roundScore int64
}

type players []player

func (ps players) Len() int           { return len(ps) }
func (ps players) Less(i, j int) bool { return ps[i].gameScore < ps[j].gameScore }
func (ps players) Swap(i, j int)      { ps[i], ps[j] = ps[j], ps[i] }
