package di

import (
	"github.com/dimayaschhu/vocabulary/pkg/cmd"
	"github.com/dimayaschhu/vocabulary/pkg/db"
	"github.com/dimayaschhu/vocabulary/pkg/httpserver"
	"github.com/dimayaschhu/vocabulary/pkg/telegram"
	"github.com/dimayaschhu/vocabulary/pkg/utils"
	"go.uber.org/fx"
)

func all() []fx.Option {
	return []fx.Option{
		db.Module,
		httpserver.Module,
		cmd.Module,
		utils.Module,
	}
}

func AppProviders() []fx.Option {
	modules := []fx.Option{
		fx.Provide(db.NewConfigProd, telegram.NewBotTest),
	}
	modules = append(modules, all()...)

	return modules
}

func TestProviders() []fx.Option {
	modules := []fx.Option{
		fx.Provide(db.NewConfigTest, telegram.NewBotTest),
	}
	modules = append(modules, all()...)

	return modules
}
