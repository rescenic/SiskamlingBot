package telegram

import (
	"log"
	"strconv"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

/*
 * Message
 */

func (c *TgContext) SendMessage(text string, chatID int64) {
	if text == "" {
		text = "Bad Request: No text supplied!"
	}

	c.TimeProc = time.Now().UTC().Sub(time.Unix(c.Message.Date, 0).UTC())
	text += "\n\n⏱ <code>" + strconv.FormatFloat(c.TimeInit.Seconds(), 'f', 3, 64) + " s</code> | ⌛ <code>" + strconv.FormatFloat(c.TimeProc.Seconds(), 'f', 3, 64) + " s</code>"

	if chatID != 0 {
		_, err := c.Bot.SendMessage(chatID, text, &gotgbot.SendMessageOpts{ParseMode: "HTML"})
		if err != nil {
			log.Println(err.Error())
			return
		}

		//c.Message = msg
		return
	}

	_, err := c.Bot.SendMessage(c.Chat.Id, text, &gotgbot.SendMessageOpts{ParseMode: "HTML"})
	if err != nil {
		log.Println(err.Error())
		return
	}

	//c.Message = msg
}

func (c *TgContext) SendMessageKeyboard(text string, chatID int64, keyb [][]gotgbot.InlineKeyboardButton) {
	if text == "" {
		text = "Bad Request: No text supplied!"
	}
	
	c.TimeProc = time.Now().UTC().Sub(time.Unix(c.Message.Date, 0).UTC())
	text += "\n\n⏱ <code>" + strconv.FormatFloat(c.TimeInit.Seconds(), 'f', 3, 64) + " s</code> | ⌛ <code>" + strconv.FormatFloat(c.TimeProc.Seconds(), 'f', 3, 64) + " s</code>"

	if chatID != 0 {
		msg, err := c.Bot.SendMessage(chatID, text, &gotgbot.SendMessageOpts{ParseMode: "HTML", ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: keyb}})
		if err != nil {
			log.Println(err.Error())
			return
		}
		
		c.Message = msg
		return
	}

	msg, err := c.Bot.SendMessage(c.Chat.Id, text, &gotgbot.SendMessageOpts{ParseMode: "HTML", ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: keyb}})
	if err != nil {
		log.Println(err.Error())
		return
	}

	c.Message = msg
}

func (c *TgContext) ReplyMessage(text string) {
	if text == "" {
		text = "Bad Request: No text supplied!"
	}
	
	c.TimeProc = time.Now().UTC().Sub(time.Unix(c.Message.Date, 0).UTC())
	text += "\n\n⏱ <code>" + strconv.FormatFloat(c.TimeInit.Seconds(), 'f', 3, 64) + " s</code> | ⌛ <code>" + strconv.FormatFloat(c.TimeProc.Seconds(), 'f', 3, 64) + " s</code>"

	msg, err := c.Context.EffectiveMessage.Reply(c.Bot, text, &gotgbot.SendMessageOpts{ParseMode: "HTML"})
	if err != nil {
		log.Println(err.Error())
		return
	}

	c.Message = msg
}

func (c *TgContext) ReplyMessageKeyboard(text string, keyb [][]gotgbot.InlineKeyboardButton) {
	if text == "" {
		text = "Bad Request: No text supplied!"
	}
	
	c.TimeProc = time.Now().UTC().Sub(time.Unix(c.Message.Date, 0).UTC())
	text += "\n\n⏱ <code>" + strconv.FormatFloat(c.TimeInit.Seconds(), 'f', 3, 64) + " s</code> | ⌛ <code>" + strconv.FormatFloat(c.TimeProc.Seconds(), 'f', 3, 64) + " s</code>"

	msg, err := c.Message.Reply(c.Bot, text, &gotgbot.SendMessageOpts{ParseMode: "HTML", ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: keyb}})
	if err != nil {
		log.Println(err.Error())
		return
	}

	c.Message = msg
}

func (c *TgContext) EditMessage(text string) {
	if text == "" {
		text = "Bad Request: No text supplied!"
	}
	
	c.TimeProc = time.Now().UTC().Sub(time.Unix(c.Message.Date, 0).UTC())
	text += "\n\n⏱ <code>" + strconv.FormatFloat(c.TimeInit.Seconds(), 'f', 3, 64) + " s</code> | ⌛ <code>" + strconv.FormatFloat(c.TimeProc.Seconds(), 'f', 3, 64) + " s</code>"

	_ , err := c.Message.EditText(c.Bot, text, &gotgbot.EditMessageTextOpts{ParseMode: "HTML"})
	if err != nil {
		log.Println(err.Error())
		return
	}

	//c.Message = msg
}

func (c *TgContext) DeleteMessage(msgId int64) {
	if msgId != 0 {
		_, err := c.Bot.DeleteMessage(c.Chat.Id, msgId)
		if err != nil {
			log.Println(err.Error())
			return
		}
		return
	}

	_, err := c.Bot.DeleteMessage(c.Chat.Id, c.Message.MessageId)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

/*
 * Callback
 */

func (c *TgContext) AnswerCallback(text string, alert bool) {
	newAnswerCallbackQueryOpts := &gotgbot.AnswerCallbackQueryOpts{
		Text:      text,
		ShowAlert: alert,
	}

	_, err := c.Callback.Answer(c.Bot, newAnswerCallbackQueryOpts)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

/*
 * ChatMember
 */

func (c *TgContext) RestrictMember(userId int64, untilDate int64) {
	if userId == 0 {
		userId = c.User.Id
	}

	if untilDate == 0 {
		untilDate = -1
	}

	newOpt := &gotgbot.RestrictChatMemberOpts{UntilDate: untilDate}
	newChatPermission := gotgbot.ChatPermissions{
		CanSendMessages:      false,
		CanSendMediaMessages: false,
		CanSendPolls:         false,
		CanSendOtherMessages: false,
	}

	_, err := c.Bot.RestrictChatMember(c.Chat.Id, userId, newChatPermission, newOpt)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
