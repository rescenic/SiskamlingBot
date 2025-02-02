package util

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func IsGroup(t string) bool {
	return t == "supergroup"
}

func IsPrivate(t string) bool {
	return t == "private"
}

func RequireGroup(b *gotgbot.Bot, ctx *ext.Context) error {
	if !IsGroup(ctx.Message.Chat.Type) {
		_, err := ctx.Message.Reply(b, "Perintah ini hanya bisa digunakan dalam grup!", nil)
		return err
	}
	return nil
}

func RequirePrivate(b *gotgbot.Bot, ctx *ext.Context) error {
	if !IsPrivate(ctx.Message.Chat.Type) {
		_, err := ctx.Message.Reply(b, "Perintah ini hanya bisa digunakan dalam japri!", nil)
		return err
	}
	return nil
}
