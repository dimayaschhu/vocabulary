package service

import (
	"github.com/dimayaschhu/vocabulary/pkg/telegram"
	"strings"
)

const CHAT_ID = 563738410

type Store interface {
	GetAsk() (string, []string)
	GetResult(answer string) bool
}

type BotService struct {
	bot   telegram.Bot
	store Store
}

func NewBotService(bot telegram.Bot, store Store) *BotService {
	return &BotService{
		store: store,
		bot:   bot,
	}
}

func (v *BotService) SendAsq() {
	ask, options := v.store.GetAsk()
	v.bot.SendMessageWithButton(CHAT_ID, ask, options)
}

func (v *BotService) SendResult(word string) {
	if v.store.GetResult(word) {
		v.bot.SendMessageWithButton(CHAT_ID, "TRUE", []string{"next"})
		return
	}

	v.bot.SendMessage(CHAT_ID, "FALSE")

}

func (v *BotService) Run() {
	updates := v.bot.GetUpdatesChan()

	for update := range updates {
		if update.Message != nil { // If we got a message
			if strings.Contains(update.Message.Text, "!") {
				v.SendResult(update.Message.Text)
				continue
			}

			v.SendAsq()
		}
	}
}
