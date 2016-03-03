package commands

import "bot"

type Info struct{}

func (Info) Help() string {
	return "bot information"
}

func (Info) Run(b *bot.Bot, cmd ...string) string {
	raw, err := b.Request(bot.METHOD_GETME, "")
	if err != nil {
		return err.Error()
	}
	return string(raw[0:])
}
