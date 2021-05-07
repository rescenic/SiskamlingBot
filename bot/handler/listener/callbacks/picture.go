package callbacks

import (
	"SiskamlingBot/bot/model"
	"context"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"log"
	"regexp"
	"strconv"
)

func PictureCB(b *gotgbot.Bot, ctx *ext.Context) error {
	cb := ctx.Update.CallbackQuery
	pattern, _ := regexp.Compile(`picture\((.+?)\)`)

	if !pattern.MatchString(cb.Data) {
		return ext.ContinueGroups
	}

	if !(pattern.FindStringSubmatch(cb.Data)[1] == strconv.Itoa(int(cb.From.Id))) {
		_, err := cb.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
			Text:      "❌ ANDA BUKAN PENGGUNA YANG DIMAKSUD!",
			ShowAlert: true,
			CacheTime: 0,
		})
		if err != nil {
			log.Println("failed to answer callbackquery: " + err.Error())
			return ext.ContinueGroups
		}
		return ext.ContinueGroups
	}

	if p, err := cb.From.GetProfilePhotos(b, nil); p != nil && p.TotalCount == 0 {
		if err != nil {
			log.Println("failed to get pictures: " + err.Error())
			return ext.ContinueGroups
		}

		_, err = cb.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
			Text:      "❌ Anda belum memasang foto profil",
			ShowAlert: true,
			CacheTime: 0,
		})
		if err != nil {
			log.Println("failed to answer callbackquery: " + err.Error())
			return ext.ContinueGroups
		}
		return ext.ContinueGroups
	}

	// Delete user status if user has set username
	err := model.DeleteUsernameByID(context.TODO(), cb.From.Id)
	if err != nil {
		log.Println("failed to save status to DB: " + err.Error())
		return ext.ContinueGroups
	}

	_, err = b.RestrictChatMember(cb.Message.Chat.Id, cb.From.Id, gotgbot.ChatPermissions{
		CanSendMessages:      true,
		CanSendMediaMessages: true,
		CanSendPolls:         true,
		CanSendOtherMessages: true,
	}, nil)
	if err != nil {
		log.Println("failed to restrict chatmember: " + err.Error())
		return ext.ContinueGroups
	}

	_, err = cb.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
		Text:      "✅ Terimakasih telah memasang Foto Profil",
		ShowAlert: true,
		CacheTime: 0,
	})
	if err != nil {
		log.Println("failed to answer callbackquery: " + err.Error())
		return ext.ContinueGroups
	}

	_, err = cb.Message.Delete(b)
	if err != nil {
		log.Println("failed to delete message: " + err.Error())
		return ext.ContinueGroups
	}

	return ext.ContinueGroups
}
