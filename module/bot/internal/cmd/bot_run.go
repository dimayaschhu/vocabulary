package cmd

import (
	"github.com/dimayaschhu/vocabulary/module/bot/internal/service"
)

type BotRunCommand struct {
	botService *service.BotService
}

func NewBotRunCommand(botService *service.BotService) *BotRunCommand {
	return &BotRunCommand{botService: botService}
}

func (c *BotRunCommand) GetName() string {
	return "bot_run"
}

func (c *BotRunCommand) Run() {
	c.botService.Run()
}
