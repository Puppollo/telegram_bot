package bot

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os/exec"
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

// run shell command cmd[0] - command, cmd[1:n] - arguments
func Run(cmd ...string) (string, error) {
	command := cmd[0]
	arguments := cmd[1:]

	out, err := exec.Command(command, arguments...).Output()
	if err != nil {
		return "", err
	}
	return string(out[:]), nil
}