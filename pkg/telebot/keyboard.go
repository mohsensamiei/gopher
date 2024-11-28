package telebot

import "github.com/mohsensamiei/gopher/v2/pkg/telegram"

type ReplyKeyboard struct {
	markup telegram.ReplyKeyboardMarkup
	remove telegram.ReplyKeyboardRemove
}

func NewReplyKeyboard() *ReplyKeyboard {
	return &ReplyKeyboard{
		markup: telegram.ReplyKeyboardMarkup{
			ResizeKeyboard:  true,
			OneTimeKeyboard: true,
		},
		remove: telegram.ReplyKeyboardRemove{
			RemoveKeyboard: true,
		},
	}
}

func (k *ReplyKeyboard) AddCol(button telegram.KeyboardButton) *ReplyKeyboard {
	if len(k.markup.Keyboard) == 0 {
		return k.AddRow(button)
	}
	k.markup.Keyboard[len(k.markup.Keyboard)-1] = append(k.markup.Keyboard[len(k.markup.Keyboard)-1], button)
	return k
}

func (k *ReplyKeyboard) AddRow(button telegram.KeyboardButton) *ReplyKeyboard {
	k.markup.Keyboard = append(k.markup.Keyboard, []telegram.KeyboardButton{button})
	return k
}

func (k *ReplyKeyboard) Render() telegram.KeyboardMarkup {
	if k.markup.Keyboard == nil {
		return k.remove
	} else {
		return k.markup
	}
}

type InlineKeyboard struct {
	markup telegram.InlineKeyboardMarkup
}

func NewInlineKeyboard() *InlineKeyboard {
	return &InlineKeyboard{
		markup: telegram.InlineKeyboardMarkup{},
	}
}

func (k *InlineKeyboard) AddCol(button telegram.InlineKeyboardButton) *InlineKeyboard {
	if len(k.markup.InlineKeyboard) == 0 {
		return k.AddRow(button)
	}
	k.markup.InlineKeyboard[len(k.markup.InlineKeyboard)-1] = append(k.markup.InlineKeyboard[len(k.markup.InlineKeyboard)-1], button)
	return k
}

func (k *InlineKeyboard) AddRow(button telegram.InlineKeyboardButton) *InlineKeyboard {
	k.markup.InlineKeyboard = append(k.markup.InlineKeyboard, []telegram.InlineKeyboardButton{button})
	return k
}

func (k *InlineKeyboard) Render() telegram.KeyboardMarkup {
	if k.markup.InlineKeyboard == nil {
		return nil
	} else {
		return k.markup
	}
}
