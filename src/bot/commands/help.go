package commands

import (
	"bot"
	"fmt"
)

type Help struct{}

func (Help) Help() string {
	return "help - this help"
}

func (Help) Run(b *bot.Bot, cmd ...string) string {
	var res string
	for name, command := range b.Commands {
		res += fmt.Sprintf("%s - %s\n", name, command.Help())
	}
	return res
}
