package fambot

import (
	"encoding/json"
	"fmt"
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
	//TODO: randomise question
	var q struct{ Questions []*question }
	err := json.Unmarshal([]byte(listofquestions), &q)
	if err != nil {
		fmt.Println("[newQuestions] Can't unmarshall listofquestions. "+
			"Err:", err.Error())
	}

	return q.Questions
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

const listofquestions = `
{
	"questions": [
		{"id": 0, "text": "Binatang yang bisa terbang selain burung?", "answers": [
			{"text": "angsa", "score": 37},
			{"text": "bebek", "score": 28},
			{"text": "ayam", "score": 22},
			{"text": "daun", "score": 17}
		]},
		{"id": 1, "text": "Hewan yang sering ditemui di jalan?", "answers": [
			{"text": "kucing", "score": 30},
			{"text": "tikus", "score": 21},
			{"text": "kecoa", "score": 14},
			{"text": "katak", "tag": ["kodok"], "score": 13},
			{"text": "laba-laba", "score": 13}
		]},
		{"id": 2, "text": "Ikan yang sering di makan?", "answers": [
			{"text": "bawal", "score": 30},
			{"text": "gurame", "score": 21},
			{"text": "lele", "score": 14},
			{"text": "cupang", "score": 13},
			{"text": "ayam-ayam", "tag": ["ayam"], "score": 13}
		]}
	]
}
`
