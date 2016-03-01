package bot

import (
	yaml "gopkg.in/yaml.v2"

	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	TELEGRAM_API_URL = "https://api.telegram.org/bot%s/%s"

	METHOD_GETME      = "getMe"
	METHOD_GETUPDATES = "getUpdates"
)

type (
	BotConfig struct {
		Token   string        `yaml:"Token"`
		Timeout time.Duration `yaml:"Timeout"`
	}

	Bot struct {
		BotConfig
		client *http.Client
		resultCounter int
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
		config = &BotConfig{Timeout: time.Second}
	}

	transport := http.Transport{
		DisableCompression:  true,
		DisableKeepAlives:   false,
		MaxIdleConnsPerHost: 1,
	}

	bot := &Bot{BotConfig: *config, resultCounter:0}
	bot.client = &http.Client{Transport: &transport}
	return bot
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
		updates, err := b.getUpdate()
		log.Println(err, updates)
		time.Sleep(b.Timeout)
	}
}
