package telegram

type EditMessageText struct {
	ChatID                int64            `json:"chat_id,omitempty"`                  // Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID             int64            `json:"message_id,omitempty"`               // Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageID       string           `json:"inline_message_id,omitempty"`        // Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	Text                  string           `json:"text"`                               // Text of the message to be sent, 1-4096 characters after entities parsing
	ParseMode             FormattingOption `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the message text. See formatting options for more details.
	Entities              []MessageEntity  `json:"entities,omitempty"`                 // Optional. A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	DisableWebPagePreview bool             `json:"disable_web_page_preview,omitempty"` // Optional. Disables link previews for links in this message
	ReplyMarkup           KeyboardMarkup   `json:"reply_markup,omitempty"`             // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// EditMessageText
// Use this method to edit text and game messages.
// On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
func (c Connection) EditMessageText(req EditMessageText) (*Message, error) {
	if req.ParseMode == "" && c.TelegramDefaultParseMode != "" {
		req.ParseMode = c.TelegramDefaultParseMode
	}
	var res Response[Message]
	if err := request(c.TelegramToken, editMessageText, req, &res); err != nil {
		return nil, err
	}
	return &res.Result, nil
}
