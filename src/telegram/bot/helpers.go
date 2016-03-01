package bot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func content(response *http.Response) ([]byte, error) {
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return contents, nil
}

func (b *Bot) request(method string, data string) (*Result, error) {
	var result Result

	request, err := http.NewRequest("POST", fmt.Sprintf(TELEGRAM_API_URL, b.Token, method), bytes.NewBuffer([]byte(data)))
	if err != nil {
		return nil, err
	}
	response, err := b.client.Do(request)
	if err != nil {
		return nil, err
	}

	content, err := content(response)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (b *Bot) getUpdate() (Updates, error) {
	result, err := b.request(METHOD_GETUPDATES, "")
	if err != nil {
		return nil, err
	}

	if !result.Ok {
		return nil, errors.New("error result.ok=false")
	}

	var updates Updates
	raw, ok := result.Result.([]byte)
	if !ok {
		return nil, errors.New("cannot convert result.result to []byte")
	}
	err = json.Unmarshal(raw, &updates)
	if err != nil {
		return nil, err
	}

	return updates, nil
}
