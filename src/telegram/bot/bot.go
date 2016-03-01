package bot

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	METHOD_GETME = "getMe"
)

type (
	BotConfig struct {
		Token   string        `yaml:"Token"`
		Timeout time.Duration `yaml:"Timeout"`
	}

	Bot struct {
		BotConfig
	}
)

func NewBotConfig(file string) (*BotConfig, error) {
	botConfig := BotConfig{}
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(raw, &botConfig)
	if err != nil {
		return nil, err
	}

	botConfig.Timeout *= time.Second

	return &botConfig, nil
}

func NewBot(config *BotConfig) *Bot {
	if config == nil {
		return &Bot{BotConfig{Timeout: time.Second}}
	}
	return &Bot{BotConfig: *config}
}

func (b *Bot) Run() {
	log.Println("token", b.Token)
	log.Println("timeout", b.Timeout)

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool)

	go func() {
		log.Println("exiting with code", <-sig)
		done <- true
	}()

	go b.handle()

	<-done
}

func (b *Bot) handle() {
	for {
		time.Sleep(b.Timeout)
		print(".")
	}
}
