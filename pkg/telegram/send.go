package telegram

type FormattingOption string

func (f FormattingOption) String() string {
	return string(f)
}

const (
	MarkdownV2 FormattingOption = "MarkdownV2"
	HTML       FormattingOption = "HTML"
	Markdown   FormattingOption = "Markdown"
)

func (f *FormattingOption) UnmarshalText(text []byte) error {
	*f = FormattingOption(text)
	return nil
}

func (f FormattingOption) MarshalText() (text []byte, err error) {
	return []byte(f), nil
}

type SendMessage struct {
	ChatID                   int              `json:"chat_id"`                               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Text                     string           `json:"text"`                                  // Text of the message to be sent, 1-4096 characters after entities parsing
	ParseMode                FormattingOption `json:"parse_mode,omitempty"`                  // Optional. Mode for parsing entities in the message text. See formatting options for more details.
	Entities                 []MessageEntity  `json:"entities,omitempty"`                    // Optional. A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	DisableWebPagePreview    bool             `json:"disable_web_page_preview,omitempty"`    // Optional. Disables link previews for links in this message
	DisableNotification      bool             `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool             `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageID         int              `json:"reply_to_message_id,omitempty"`         // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              KeyboardMarkup   `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// SendMessage
// Use this method to send text messages. On success, the sent Message is returned.
// reply_markup	InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
func (c Connection) SendMessage(req SendMessage) (*Message, error) {
	if req.ParseMode == "" && c.TelegramDefaultParseMode != "" {
		req.ParseMode = c.TelegramDefaultParseMode
	}
	var res Response[Message]
	if err := request(c.TelegramToken, sendMessage, req, &res); err != nil {
		return nil, err
	}
	return &res.Result, nil
}

type Action string

func (f Action) String() string {
	return string(f)
}

const (
	Typing          Action = "typing"
	UploadPhoto     Action = "upload_photo"
	RecordVideo     Action = "record_video"
	UploadVideo     Action = "upload_video"
	RecordVoice     Action = "record_voice"
	UploadVoice     Action = "upload_voice"
	UploadDocument  Action = "upload_document"
	ChooseSticker   Action = "choose_sticker"
	FindLocation    Action = "find_location"
	RecordVideoNote Action = "record_video_note"
	UploadVideoNote Action = "upload_video_note"
)

type SendAction struct {
	ChatID          int    `json:"chat_id"`                     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Action          Action `json:"action"`                      // Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_voice or upload_voice for voice notes, upload_document for general files, choose_sticker for stickers, find_location for location data, record_video_note or upload_video_note for video notes.
	MessageThreadID int    `json:"message_thread_id,omitempty"` // Optional. Unique identifier for the target message thread; supergroups only
}

// SendAction
// Use this method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.
// Example: The ImageBot needs some time to process a request and upload the image. Instead of sending a text message along the lines of “Retrieving image, please wait…”, the bot may use sendChatAction with action = upload_photo. The user will see a “sending photo” status for the bot.
// We only recommend using this method when a response from the bot will take a noticeable amount of time to arrive.
func (c Connection) SendAction(req SendAction) (bool, error) {
	var res Response[bool]
	if err := request(c.TelegramToken, sendChatAction, req, &res); err != nil {
		return false, err
	}
	return res.Result, nil
}

type InputFileID struct {
	ID string
}

func (i InputFileContent) InputFileID() {

}

type InputFileContent struct {
	Name    string
	Content []byte
}

func (i InputFileContent) isInputFile() {

}

type InputFile interface {
	isInputFile()
}

type SendDocument struct {
	ChatID                      int              `json:"chat_id"`                                  // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID             int              `json:"message_thread_id,omitempty"`              // Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Document                    InputFile        `json:"document"`                                 // File to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data
	Thumbnail                   InputFile        `json:"thumbnail,omitempty"`                      // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>
	Caption                     string           `json:"caption"`                                  // Text of the message to be sent, 1-4096 characters after entities parsing
	ParseMode                   FormattingOption `json:"parse_mode,omitempty"`                     // Optional. Mode for parsing entities in the message text. See formatting options for more details.
	CaptionEntities             []MessageEntity  `json:"caption_entities,omitempty"`               // Optional. A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	DisableContentTypeDetection bool             `json:"disable_content_type_detection,omitempty"` // Optional. Disables automatic server-side content type detection for files uploaded using multipart/form-data
	DisableNotification         bool             `json:"disable_notification,omitempty"`           // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent              bool             `json:"protect_content,omitempty"`                // Optional. Protects the contents of the sent message from forwarding and saving
	ReplyToMessageID            int              `json:"reply_to_message_id,omitempty"`            // Optional. If the message is a reply, ID of the original message
	AllowSendingWithoutReply    bool             `json:"allow_sending_without_reply,omitempty"`    // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup                 KeyboardMarkup   `json:"reply_markup,omitempty"`                   // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// SendDocument
// Use this method to send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
func (c Connection) SendDocument(req SendDocument) (*Message, error) {
	if req.ParseMode == "" && c.TelegramDefaultParseMode != "" {
		req.ParseMode = c.TelegramDefaultParseMode
	}
	var res Response[Message]
	if err := request(c.TelegramToken, sendDocument, req, &res); err != nil {
		return nil, err
	}
	return &res.Result, nil
}
