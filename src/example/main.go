package main

import (
	"bot"
	"bot/commands"
	"flag"
	"log"
)

func main() {
	file := flag.String("c", "config.yaml", "-c config.yaml")
	flag.Parse()

	botConfig, err := bot.NewBotConfig(*file)
	if err != nil {
		log.Println(err.Error())
	}

	bot := bot.NewBot(botConfig, bot.Commands{
		"help": commands.Help{},
		"list": commands.List{},
		"run":  commands.Run{},
		"info": commands.Info{},
	})

	bot.Run()
}
