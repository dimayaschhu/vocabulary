package bot

import (
	"github.com/dimayaschhu/vocabulary/module/bot/internal/cmd"
	"github.com/dimayaschhu/vocabulary/module/bot/internal/repo"
	"github.com/dimayaschhu/vocabulary/module/bot/internal/service"
	cmd2 "github.com/dimayaschhu/vocabulary/pkg/cmd"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		repo.NewWordRepo,
		service.NewBotService,
		service.NewStoreService,
		cmd.NewBotRunCommand,
	),
	fx.Invoke(func(command *cmd.BotRunCommand, register *cmd2.Register) {
		register.Add(command)
	}),
)
