package tests

import (
	"context"
	"flag"
	"github.com/dimayaschhu/vocabulary/module/bot"
	"github.com/dimayaschhu/vocabulary/module/web"
	"github.com/dimayaschhu/vocabulary/pkg/di"
	"github.com/dimayaschhu/vocabulary/pkg/httpserver"
	"github.com/dimayaschhu/vocabulary/tests/steps"
	"os"
	"path"
	"runtime"
	"testing"
	"time"

	"github.com/OpenTagOS/allure2-godog/allure"
	"github.com/OpenTagOS/allure2-godog/alluregodog"

	"github.com/pkg/errors"

	"go.uber.org/fx"

	"github.com/cucumber/godog"
)

var tags = flag.String("godog.tags", "", "")

// Switch working dir to project root (by default "go test ./..." set it to ./tests/).
func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func TestAllSuites(test *testing.T) {
	allureWriter := allure.NewReportWriter("./allure/")
	godog.Format("allure", "Allure 2 formatter", alluregodog.NewFormatter(allureWriter))

	opts := godog.Options{
		Format:    "progress", // "progress|allure"
		Strict:    true,
		Paths:     []string{"./tests/features"},
		Randomize: time.Now().UTC().UnixNano(), // randomize scenario execution order
		Tags:      *tags,
	}

	fxOptions := di.TestProviders()
	fxOptions = append(fxOptions, bot.Module)
	fxOptions = append(fxOptions, web.Module)
	fxOptions = append(fxOptions, fx.Invoke(func(server *httpserver.Router) {
		go server.GetEngine().Run()
	}))

	fxOptions = append(
		fxOptions,
		fx.Provide(
			steps.NewCmdStepHandler,
			steps.NewDBStepHandler,
			steps.NewHTTPStepHandler,
			steps.NewBotStepHandler,
		),

		fx.Invoke(func(
			lc fx.Lifecycle,
			dbStepsHandler *steps.DBStepHandler,
			cmdStepHandler *steps.CmdStepHandler,
			httpStepHandler *steps.HTTPStepHandler,
			botStepHandler *steps.BotStepHandler,
		) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {

					suite := godog.TestSuite{
						Name: "backend",
						ScenarioInitializer: func(ctx *godog.ScenarioContext) {
							ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
								println("Before------------")
								return ctx, nil
							})

							ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
								dbStepsHandler.RemoveDB()
								return ctx, nil
							})

							dbStepsHandler.RegisterSteps(ctx)
							cmdStepHandler.RegisterSteps(ctx)
							httpStepHandler.RegisterSteps(ctx)
							botStepHandler.RegisterSteps(ctx)
						},
						Options: &opts,
					}

					status := suite.Run()
					if status != 0 {
						return errors.New("func tests failed")
					}

					return nil
				},
			})
		}),
	)

	app := fx.New(fxOptions...)

	if err := app.Start(context.Background()); err != nil {
		test.Error(err)
	}
}
