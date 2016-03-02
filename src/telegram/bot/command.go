package bot

import (
	"encoding/json"
	"errors"
	"strings"
)

const (
	RUN = "run"
)

func (b *Bot) handleCommand(update *Update) *SendMessageRequest {
	if update == nil {
		return nil
	}

	command := strings.Fields(update.Message.Text)
	response := SendMessageRequest{Id:update.Message.Chat.Id}

	switch command[0] {
	case RUN:
		output, err := Run(command[1:]...)
		if err != nil {
			response.Text = err.Error()
		}else {
			response.Text = output
		}
	}

	return &response
}

func (b *Bot) getUpdate(request GetUpdateRequest) ([]Update, error) {
	var updateResult UpdateResult

	raw, err := b.request(METHOD_GETUPDATES, request.String())
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(raw, &updateResult)
	if err != nil {
		return nil, err
	}

	if !updateResult.Ok {
		return nil, errors.New("error result.ok=false")
	}

	return updateResult.Result, nil
}

func (b *Bot) sendMessage(request *SendMessageRequest) ([]byte, error) {
	if request == nil {
		return nil, nil
	}
	raw, err := b.request(METHOD_SENDMESSAGE, request.String())
	if err != nil {
		return nil, err
	}
	return raw, nil
}