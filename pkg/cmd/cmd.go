package cmd

import "go.uber.org/fx"

type Command interface {
	GetName() string
	Run()
}

type Register struct {
	commands []Command
}

func (r *Register) Add(command Command) {
	r.commands = append(r.commands, command)
}

func (r *Register) Run(nameCommand string) {
	for _, c := range r.commands {
		if c.GetName() == nameCommand {
			c.Run()
		}

	}
}

func NewRegister() *Register {
	return &Register{}
}

var Module = fx.Provide(NewRegister)
