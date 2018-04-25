package fambot

import (
	"math/rand"
	"strings"
)

type question struct {
	QuestionID int64     `json:"id"`
	Text       string    `json:"text"`
	Answers    []*answer `json:"answers"`
}

type answer struct {
	Text   string   `json:"text"`    // the matching answer.
	Tag    []string `json:"tag"`     // tag contains similar word to text.
	UserID string   `json:"user_id"` // user who answer the question correctly.
	Score  int64    `json:"score"`   // score of the answer.
}

func (s *server) newQuestions() []*question {
	//TODO: created newQuestions from database instead from text file.
	//TODO: generate DB in init instead of parsing here everytime.

	// var q struct{ Questions []*question }
	// err := json.Unmarshal([]byte(listofquestions), &q)
	// if err != nil {
	// 	fmt.Println("[newQuestions] Can't unmarshall listofquestions. "+
	// 		"Err:", err.Error())
	// }

	var qs = make([]*question, 0)
	var qsStr = strings.Split(strings.TrimSpace(listofquestions), ";")
	for i := 0; i < len(qsStr); i++ {
		var q question
		var asStr = strings.Split(strings.TrimSpace(qsStr[i]), "|")
		if len(asStr) == 6 {
			var as = make([]*answer, 0)
			var scr int64 = 32
			for j := 1; j <= 5; j++ {
				var a answer
				var ts = strings.Split(strings.ToLower(strings.TrimSpace(asStr[j])), "~")
				if len(ts) > 0 {
					a = answer{
						Text:  ts[0],
						Tag:   ts,
						Score: scr,
					}
					as = append(as, &a)
				}
				scr -= 6
			}
			q = question{
				QuestionID: int64(i),
				Text:       asStr[0],
				Answers:    as,
			}
			qs = append(qs, &q)
		}
	}

	for i := range qs {
		j := rand.Intn(i + 1)
		qs[i], qs[j] = qs[j], qs[i]
	}

	return qs
}

func (q *question) answer(userID, answer string) *answer {
	for i := range q.Answers {
		// check if it isn't answered before
		if q.Answers[i].UserID == "" {
			var answers = q.Answers[i].Tag
			answers = append(answers, q.Answers[i].Text)
			for j := range answers {
				// TODO: enhanced using another comparison method
				if answers[j] == answer {
					q.Answers[i].UserID = userID
					return q.Answers[i]
				}
			}
		}
	}
	return nil
}
