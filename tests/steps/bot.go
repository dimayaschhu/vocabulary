package steps

import (
	"github.com/cucumber/godog"
	"github.com/dimayaschhu/vocabulary/pkg/telegram"
)

type BotStepHandler struct {
	bot telegram.Bot
}

func NewBotStepHandler(bot telegram.Bot) *BotStepHandler {
	return &BotStepHandler{
		bot: bot,
	}
}

func (h *BotStepHandler) RegisterSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^pull message "([^"]*)"`, h.pullMessage)
}

func (h *BotStepHandler) pullMessage(msg string) error {
	h.bot.PullMessage(msg)
	return nil
}
