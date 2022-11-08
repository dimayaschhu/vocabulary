package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot interface {
	SendMessageWithButton(chatId int64, msg string, buttons []string)
	SendMessage(chatId int64, msg string)
	GetUpdatesChan() tgbotapi.UpdatesChannel
	PullMessage(msg string)
}

type TelegramBot struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramBot() Bot {
	return &TelegramBot{}
}

func getBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI("5673296847:AAE2Z2Bz1uFt1SVL6O3r3khshGBbAnbbHeg")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	return bot
}

func (t *TelegramBot) SendMessageWithButton(chatId int64, text string, buttons []string) {
	t.bot = getBot()
	var k []tgbotapi.KeyboardButton
	for _, button := range buttons {
		k = append(k, tgbotapi.NewKeyboardButton(button))
	}

	numericKeyboard := tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(k...))

	msg := tgbotapi.NewMessage(chatId, text)
	msg.ReplyMarkup = numericKeyboard
	t.bot.Send(msg)
}

func (t *TelegramBot) SendMessage(chatId int64, text string) {
	t.bot = getBot()
	msg := tgbotapi.NewMessage(chatId, text)
	t.bot.Send(msg)
}

func (t *TelegramBot) GetUpdatesChan() tgbotapi.UpdatesChannel {
	t.bot = getBot()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return t.bot.GetUpdatesChan(u)
}

func (t *TelegramBot) PullMessage(msg string) {}

type BotTest struct {
	c chan tgbotapi.Update
}

func NewBotTest() Bot {
	return &BotTest{c: make(chan tgbotapi.Update)}
}

func (t *BotTest) SendMessageWithButton(chatId int64, msg string, buttons []string) {

}

func (t *BotTest) SendMessage(chatId int64, msg string) {

}

func (t *BotTest) GetUpdatesChan() tgbotapi.UpdatesChannel {
	return t.c
}

func (t *BotTest) PullMessage(data string) {
	var msg = tgbotapi.Update{
		UpdateID: 1,
		Message: &tgbotapi.Message{
			Text: data,
			From: &tgbotapi.User{ID: 234},
		},
	}
	go func() {
		t.c <- msg
	}()
}
