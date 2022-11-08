package service

import (
	"github.com/dimayaschhu/vocabulary/module/bot/internal/repo"
	"strings"
)

type StoreService struct {
	wordRepo *repo.WordRepo
	history  []string
	words    map[string]string
}

func NewStoreService(wordRepo *repo.WordRepo) Store {
	return &StoreService{wordRepo: wordRepo, words: make(map[string]string)}
}

func (s *StoreService) GetAsk() (string, []string) {
	words, err := s.wordRepo.GetWordsByLimit(3)
	if err != nil {
		println("-----------")
		println(err.Error())
		println("-----------")
	}
	w := words[0]
	ask := w.Name + "?"
	s.history = append(s.history, ask)
	s.words[ask] = w.Translate + "!"

	f1 := words[1]

	f2 := words[2]

	s.history = append(s.history, s.words[ask]+"|"+f1.Translate+"!|"+f2.Translate+"!")

	return ask, []string{s.words[ask], f1.Translate + "!", f2.Translate + "!"}
}

func (s *StoreService) GetResult(answer string) bool {
	var ask string

	for _, r := range s.history {
		if strings.Contains(r, answer) {
			continue
		}
		ask = r
	}

	res := s.words[ask]

	return res == answer
}
