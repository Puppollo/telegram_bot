package bot

import "encoding/json"

type (
	UpdateResult struct {
		Ok     bool	`json:"ok"`
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
)

func (u *UpdateResult) String() string {
	raw, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(raw)
}