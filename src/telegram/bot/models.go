package bot

import (
	"encoding/json"
)

type (
	UpdateResult struct {
		Ok     bool     `json:"ok"`
		Result []Update `json:"result"`
	}

	Update struct {
		Id      int     `json:"update_id"`
		Message Message `json:"message"`
	}

	Message struct {
		Id   int    `json:"message_id"` //Unique message identifier
		From User   `json:"from"`       //Optional. Sender, can be empty for messages sent to channels
		Date int    `json:"date"`       //Date the message was sent in Unix time
		Chat Chat   `json:"chat"`       //Conversation the message belongs to
		Text string `json:"text"`       //Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters.
	}

	Chat struct {
		Id        int    `json:"id"`         //	Unique identifier for this chat, not exceeding 1e13 by absolute value
		Type      string `json:"type"`       //	Type of chat, can be either “private”, “group”, “supergroup” or “channel”
		Title     string `json:"title"`      //	Optional. Title, for channels and group chats
		Username  string `json:"username"`   // Optional. Username, for private chats and channels if available
		FirstName string `json:"first_name"` //	Optional. First name of the other party in a private chat
		LastName  string `json:"last_name"`  //	Optional. Last name of the other party in a private chat
	}

	User struct {
		Id        int    `json:"id"`         //Unique identifier for this user or bot
		FirstName string `json:"first_name"` //User‘s or bot’s first name
		LastName  string `json:"last_name"`  //Optional. User‘s or bot’s last name
		Username  string `json:"username"`   //Optional. User‘s or bot’s username
	}

	GetUpdateRequest struct {
		Offset  int `json:"offset"`  // Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue. All previous updates will forgotten.
		Limit   int `json:"limit"`   // Limits the number of updates to be retrieved. Values between 1—100 are accepted. Defaults to 100.
		Timeout int `json:"timeout"` // Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling
	}
)

func (u *UpdateResult) String() string {
	raw, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(raw)
}

func (u *GetUpdateRequest) String() string {
	raw, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(raw)
}
