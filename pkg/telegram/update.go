package telegram

import (
	"strings"
)

type UpdateType string

func (f UpdateType) String() string {
	return string(f)
}

const (
	MessageUpdate            UpdateType = "message"
	EditedMessageUpdate      UpdateType = "edited_message"
	ChannelPostUpdate        UpdateType = "channel_post"
	EditedChannelPostUpdate  UpdateType = "edited_channel_post"
	InlineQueryUpdate        UpdateType = "inline_query"
	ChosenInlineResultUpdate UpdateType = "chosen_inline_result"
	CallbackQueryUpdate      UpdateType = "callback_query"
	ShippingQueryUpdate      UpdateType = "shipping_query"
	PreCheckoutQueryUpdate   UpdateType = "pre_checkout_query"
	PollUpdate               UpdateType = "poll"
	PollAnswerUpdate         UpdateType = "poll_answer"
	MyChatMemberUpdate       UpdateType = "my_chat_member"
	ChatMemberUpdate         UpdateType = "chat_member"
	ChatJoinRequestUpdate    UpdateType = "chat_join_request"
)

// Update
// This object represents an incoming update.
// At most one of the optional parameters can be present in any given update.
type Update struct {
	UpdateID           uint                `json:"update_id"`            // The update's unique identifier. Update identifiers start from a certain positive number and increase sequentially. This ID becomes especially handy if you're using webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order. If there are no new updates for at least a week, then identifier of the next update will be chosen randomly instead of sequentially.
	Message            *Message            `json:"message"`              // Optional. New incoming message of any kind - text, photo, sticker, etc.
	EditedMessage      *Message            `json:"edited_message"`       // Optional. New version of a message that is known to the bot and was edited
	ChannelPost        *Message            `json:"channel_post"`         // Optional. New incoming channel post of any kind - text, photo, sticker, etc.
	EditedChannelPost  *Message            `json:"edited_channel_post"`  // Optional. New version of a channel post that is known to the bot and was edited
	InlineQuery        *InlineQuery        `json:"inline_query"`         // Optional. New incoming inline query
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"` // Optional. The result of an inline query that was chosen by a user and sent to their chat partner. Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
	CallbackQuery      *CallbackQuery      `json:"callback_query"`       // Optional. New incoming callback query
	ShippingQuery      *ShippingQuery      `json:"shipping_query"`       // Optional. New incoming shipping query. Only for invoices with flexible price
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query"`   // Optional. New incoming pre-checkout query. Contains full information about checkout
	Poll               *Poll               `json:"poll"`                 // Optional. New poll state. Bots receive only updates about stopped polls and polls, which are sent by the bot
	PollAnswer         *PollAnswer         `json:"poll_answer"`          // Optional. A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were sent by the bot itself.
	MyChatMember       *ChatMemberUpdated  `json:"my_chat_member"`       // Optional. The bot's chat member status was updated in a chat. For private chats, this update is received only when the bot is blocked or unblocked by the user.
	ChatMember         *ChatMemberUpdated  `json:"chat_member"`          // Optional. A chat member's status was updated in a chat. The bot must be an administrator in the chat and must explicitly specify “chat_member” in the list of allowed_updates to receive these updates.
	ChatJoinRequest    *ChatJoinRequest    `json:"chat_join_request"`    // Optional. A request to join the chat has been sent. The bot must have the can_invite_users administrator right in the chat to receive these updates.
}

func (u Update) Type() UpdateType {
	if u.Message != nil {
		return MessageUpdate
	}
	if u.EditedMessage != nil {
		return EditedMessageUpdate
	}
	if u.ChannelPost != nil {
		return ChannelPostUpdate
	}
	if u.EditedChannelPost != nil {
		return EditedChannelPostUpdate
	}
	if u.InlineQuery != nil {
		return InlineQueryUpdate
	}
	if u.ChosenInlineResult != nil {
		return ChosenInlineResultUpdate
	}
	if u.CallbackQuery != nil {
		return CallbackQueryUpdate
	}
	if u.ShippingQuery != nil {
		return ShippingQueryUpdate
	}
	if u.PreCheckoutQuery != nil {
		return PreCheckoutQueryUpdate
	}
	if u.Poll != nil {
		return PollUpdate
	}
	if u.PollAnswer != nil {
		return PollAnswerUpdate
	}
	if u.MyChatMember != nil {
		return MyChatMemberUpdate
	}
	if u.ChatMember != nil {
		return ChatMemberUpdate
	}
	if u.ChatJoinRequest != nil {
		return ChatJoinRequestUpdate
	}
	return ""
}

func (u Update) From() *User {
	if u.Message != nil {
		return u.Message.From
	}
	if u.CallbackQuery != nil {
		return &u.CallbackQuery.From
	}
	return nil
}

func (u Update) ChatID() int {
	if u.Message != nil {
		return u.Message.Chat.ID
	}
	if u.CallbackQuery != nil {
		return u.CallbackQuery.Message.Chat.ID
	}
	return 0
}

func (u Update) MessageID() int {
	if u.Message != nil {
		return u.Message.MessageID
	}
	if u.CallbackQuery != nil {
		return u.CallbackQuery.Message.MessageID
	}
	return 0
}

func (u Update) IsCommand() bool {
	if u.Type() != MessageUpdate {
		return false
	}
	for _, entity := range u.Message.Entities {
		if entity.Type == BotCommandEntity {
			return true
		}
	}
	return false
}

func (u Update) IsSimilarCommand(commands []string) bool {
	if u.Type() != MessageUpdate {
		return false
	}
	for _, command := range commands {
		if command == strings.ToLower(strings.TrimSpace(u.Message.Text)) {
			return true
		}
	}
	return false
}
