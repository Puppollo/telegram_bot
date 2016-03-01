package bot

type (
	Update struct {
		Update_id int
		Message   Message
	}

	Message struct {
		message_id   int  //Unique message identifier
		from         User //Optional. Sender, can be empty for messages sent to channels
		date         int  //Date the message was sent in Unix time
		chat         Chat //Conversation the message belongs to
		forward_from User //Optional. For forwarded messages, sender of the original message
		forward_date int  //Optional. For forwarded messages, date the original message was sent in Unix time
		//reply_to_message MessageObj    //Optional. For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
		text string //Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters.
		//audio	Audio	//Optional. Message is an audio file, information about the file
		//document	Document	//Optional. Message is a general file, information about the file
		//photo	Array of PhotoSize	//Optional. Message is a photo, available sizes of the photo
		//sticker	Sticker	//Optional. Message is a sticker, information about the sticker
		//video	Video	//Optional. Message is a video, information about the video
		//voice	Voice	//Optional. Message is a voice message, information about the file
		//caption	string	//Optional. Caption for the photo or video, 0-200 characters
		//contact	Contact	//Optional. Message is a shared contact, information about the contact
		//location	Location	//Optional. Message is a shared location, information about the location
		new_chat_participant  User   //Optional. A new member was added to the group, information about them (this member may be the bot itself)
		left_chat_participant User   //Optional. A member was removed from the group, information about them (this member may be the bot itself)
		new_chat_title        string //Optional. A chat title was changed to this value
		//new_chat_photo	Array of PhotoSize	//Optional. A chat photo was change to this value
		//delete_chat_photo	True	//Optional. Service message: the chat photo was deleted
		group_chat_created      bool //Optional. Service message: the group has been created
		supergroup_chat_created bool //Optional. Service message: the supergroup has been created
		channel_chat_created    bool //Optional. Service message: the channel has been created
		migrate_to_chat_id      int  //Optional. The group has been migrated to a supergroup with the specified identifier, not exceeding 1e13 by absolute value
		migrate_from_chat_id    int  //Optional. The supergroup has been migrated from a group with the specified identifier, not exceeding 1e13 by absolute value
	}

	Chat struct {
		id         int    //	Unique identifier for this chat, not exceeding 1e13 by absolute value
		Type       string //	Type of chat, can be either “private”, “group”, “supergroup” or “channel”
		Title      string //	Optional. Title, for channels and group chats
		Username   string // Optional. Username, for private chats and channels if available
		First_name string //	Optional. First name of the other party in a private chat
		Last_name  string //	Optional. Last name of the other party in a private chat
	}

	User struct {
		Id         int    //Unique identifier for this user or bot
		First_name string //User‘s or bot’s first name
		Last_name  string //Optional. User‘s or bot’s last name
		Username   string //Optional. User‘s or bot’s username
	}
)
