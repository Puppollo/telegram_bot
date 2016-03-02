package bot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func content(response *http.Response) ([]byte, error) {
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return contents, nil
}

func (b *Bot) request(method string, data string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf(TELEGRAM_API_URL, b.Token, method), bytes.NewBuffer([]byte(data)))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := b.client.Do(request)
	if err != nil {
		return nil, err
	}

	content, err := content(response)
	if err != nil {
		return nil, err
	}

	return content, nil
}

// check if user in list
func (b *Bot) validUser(id int) bool {
	for _, userId := range b.Users {
		if userId == id {
			return true
		}
	}
	return false
}

func (b *Bot) run(cmd ...string) (string, error) {

	if len(cmd) == 0 {
		return "", errors.New("no command")
	}

	command := fmt.Sprint(b.Commands, "/", cmd[0])
	arguments := cmd[1:]

	if _, err := os.Stat(command); err != nil {
		return "", errors.New("command not exist")
	}

	out, err := exec.Command(command, arguments...).Output()
	if err != nil {
		return "", err
	}
	return string(out[:]), nil
}

func (b *Bot) handleCommand(update *Update) *SendMessageRequest {
	if update == nil {
		return nil
	}

	command := strings.Fields(update.Message.Text)
	response := SendMessageRequest{Id: update.Message.Chat.Id, Text: "undefined command"}

	if output, ok := b.commands[command[0]]; ok {
		response.Text = output(b, command[1:]...)
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
