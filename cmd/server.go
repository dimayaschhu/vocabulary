package main

import (
	"context"
	"github.com/dimayaschhu/vocabulary/module/web"
	"github.com/dimayaschhu/vocabulary/pkg/di"
	"github.com/dimayaschhu/vocabulary/pkg/httpserver"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func NewServer() *cobra.Command {
	migrationsCmd := &cobra.Command{
		Use: "server",
	}

	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Execute migrations",
		RunE: func(cmd *cobra.Command, args []string) error {
			serverMain()
			return nil
		},
	}

	migrationsCmd.AddCommand(runCmd)

	return migrationsCmd
}

func serverMain() {
	fxOptions := di.AppProviders()
	fxOptions = append(fxOptions, web.Module)
	fxOptions = append(fxOptions, fx.Invoke(func(server *httpserver.Router) {
		server.GetEngine().Run()
	}))

	app := fx.New(fxOptions...)

	if err := app.Start(context.Background()); err != nil {
		panic(err)
	}
}
