package bot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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
