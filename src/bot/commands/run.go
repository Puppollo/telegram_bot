package commands

import "bot"

type Run struct{}

func (Run) Help() string {
	return "run <command> [<arguments>] - execute command"
}

func (Run) Run(b *bot.Bot, cmd ...string) string {
	output, err := b.Execute(cmd...)
	if err != nil {
		return err.Error()
	}
	return output
}
