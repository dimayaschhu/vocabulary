package main

import (
	"context"
	"github.com/dimayaschhu/vocabulary/module/bot"
	cmd2 "github.com/dimayaschhu/vocabulary/pkg/cmd"
	"github.com/dimayaschhu/vocabulary/pkg/di"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func NewBotCommand() *cobra.Command {
	migrationsCmd := &cobra.Command{
		Use: "bot",
	}

	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Execute migrations",
		RunE: func(cmd *cobra.Command, args []string) error {
			botMain()
			return nil
		},
	}

	migrationsCmd.AddCommand(runCmd)

	return migrationsCmd
}

func botMain() {
	fxOptions := di.AppProviders()
	fxOptions = append(fxOptions, bot.Module)
	fxOptions = append(fxOptions, fx.Invoke(func(register *cmd2.Register) {
		register.Run("bot_run")
	}))

	app := fx.New(fxOptions...)

	if err := app.Start(context.Background()); err != nil {
		panic(err)
	}
}
