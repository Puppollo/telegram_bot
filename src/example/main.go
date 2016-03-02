package main

import (
	"bot"
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

	bot := bot.NewBot(botConfig)
	bot.Run()
}