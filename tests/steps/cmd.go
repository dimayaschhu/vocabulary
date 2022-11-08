package steps

import (
	"github.com/cucumber/godog"
	"github.com/dimayaschhu/vocabulary/pkg/cmd"
)

type CmdStepHandler struct {
	register *cmd.Register
}

func NewCmdStepHandler(register *cmd.Register) *CmdStepHandler {
	return &CmdStepHandler{
		register: register,
	}
}

func (h *CmdStepHandler) RegisterSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^execute command "([^"]*)"`, h.executeCommand)
}

func (h *CmdStepHandler) executeCommand(name string) error {
	println("executeCommand")
	h.register.Run(name)
	return nil
}
