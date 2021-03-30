package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var prob = []bool{true, true, false, false, false}

var curRand *rand.Rand

var messages map[int64]map[string]*tgbotapi.Message

func init() {
	curRand = rand.New(rand.NewSource(time.Now().Unix()))
	messages = make(map[int64]map[string]*tgbotapi.Message)
}

func main() {
	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	bot, err := tgbotapi.NewBotAPI(telegramToken)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			var gif tgbotapi.AnimationConfig
			switch update.Message.Command() {
			case "durr":
				gif = tgbotapi.NewAnimationShare(update.Message.Chat.ID, "CgACAgQAAxkBAAMdYGDU-aGj-BAjTa0dGviX35Z0gP4AAjYCAALYL5VSHxQhj1TU2cYeBA")
				if message, ok := messages[update.Message.Chat.ID][strings.ToLower(update.Message.CommandArguments())]; !ok {
					gif.Caption = randomCase(fmt.Sprintf("hurr durr %s, %s has no messages in here yet", update.Message.From.FirstName, strings.ToTitle(update.Message.CommandArguments())))
				} else {
					gif.Caption = randomCase(message.Text)
					gif.ReplyToMessageID = message.MessageID
				}
			case "fuckyou":
				gif = tgbotapi.NewAnimationShare(update.Message.Chat.ID, "CgACAgQAAxkBAAIBGGBhJ_lZyDk1_YmahLEsqFZ1ON9MAAJxAgACeqeMUjwXqK1QN7qGHgQ")
				if message, ok := messages[update.Message.Chat.ID][strings.ToLower(update.Message.CommandArguments())]; !ok {
					gif.Caption = fmt.Sprintf("Hey, hey %s. Fuck you, no messages from %s yet", update.Message.From.FirstName, strings.ToTitle(update.Message.CommandArguments()))
				} else {
					gif.ReplyToMessageID = message.MessageID
				}
			case "sploosh":
				gif = tgbotapi.NewAnimationShare(update.Message.Chat.ID, "CgACAgQAAxkBAAIBVWBhLj-iPzQzDDiwNi75UT7bTKk2AAKDAgACg0mNUnjNm9ca7IUmHgQ")
				if message, ok := messages[update.Message.Chat.ID][strings.ToLower(update.Message.CommandArguments())]; !ok {
					gif.Caption = fmt.Sprintf("You splooshed for nothing %s, %s has no messages in here yet", update.Message.From.FirstName, strings.ToTitle(update.Message.CommandArguments()))
				} else {
					gif.ReplyToMessageID = message.MessageID
				}
			case "goldstar":
				gif = tgbotapi.NewAnimationShare(update.Message.Chat.ID, "CgACAgQAAxkBAAICOGBiJptsJgqu7CjI7wglyFpry5rFAAJUAgACu-uVUt2pTnE70uewHgQ")
				if message, ok := messages[update.Message.Chat.ID][strings.ToLower(update.Message.CommandArguments())]; !ok {
					gif.Caption = fmt.Sprintf("%s gets a dumbass gold star for no name sent, what a dumbass!", update.Message.From.FirstName)
				} else {
					gif.ReplyToMessageID = message.MessageID
				}
			case "magma":
				gif = tgbotapi.NewAnimationShare(update.Message.Chat.ID, "CgACAgEAAxkBAAIEcmBicgTC9R0ZMNginJB6ewKudpfhAAJsAQACXyUQR0jwan3HzlslHgQ")
				if message, ok := messages[update.Message.Chat.ID][strings.ToLower(update.Message.CommandArguments())]; !ok {
					gif.Caption = fmt.Sprintf("Not fucking lava dude %s", update.Message.From.FirstName)
				} else {
					gif.ReplyToMessageID = message.MessageID
				}
			case "boobs":
				gif = tgbotapi.NewAnimationShare(update.Message.Chat.ID, "CgACAgQAAxkBAAIEpmBieuj2CmF42NGZFMj5-LViv7sNAAIqAgACno3sUSMPJOU4zQYYHgQ")
				if message, ok := messages[update.Message.Chat.ID][strings.ToLower(update.Message.CommandArguments())]; !ok {
					gif.Caption = fmt.Sprintf("%s thirsty?", update.Message.From.FirstName)
				} else {
					gif.ReplyToMessageID = message.MessageID
				}
			case "npc":
				gif = tgbotapi.NewAnimationShare(update.Message.Chat.ID, "CgACAgQAAxkBAAIEsGBie58bKOeSXEHrVrvh-Gta5Wk_AAIrAgACiX-VUpZXPFLOB0VfHgQ")
				if message, ok := messages[update.Message.Chat.ID][strings.ToLower(update.Message.CommandArguments())]; !ok {
					gif.Caption = fmt.Sprintf("%s, burn an inn lately?", update.Message.From.FirstName)
				} else {
					gif.ReplyToMessageID = message.MessageID
				}
			case "nerd":
				gif = tgbotapi.NewAnimationShare(update.Message.Chat.ID, "CgACAgQAAxkBAAIEs2BifFqTRTFAL-cM4Bhv66NW-FBoAAIaAgACP7CcUooP90RkGHy0HgQ")
				if message, ok := messages[update.Message.Chat.ID][strings.ToLower(update.Message.CommandArguments())]; !ok {
					gif.Caption = fmt.Sprintf("%s, you're a fucking nerd!", update.Message.From.FirstName)
				} else {
					gif.ReplyToMessageID = message.MessageID
				}
			case "wtf":
				gif = tgbotapi.NewAnimationShare(update.Message.Chat.ID, "CgACAgQAAxkBAAICbWBiMQW6JtqhSaQRUI3s_EXSI1TuAAJJAgACiAOUUrtigLpU3wfdHgQ")
				if message, ok := messages[update.Message.Chat.ID][strings.ToLower(update.Message.CommandArguments())]; !ok {
					gif.Caption = fmt.Sprintf("Not event an actual username WTF %s", update.Message.From.FirstName)
				} else {
					gif.ReplyToMessageID = message.MessageID
				}
			default:
				continue
			}
			_, _ = bot.Send(gif)
		} else {
			if messages[update.Message.Chat.ID] == nil {
				messages[update.Message.Chat.ID] = make(map[string]*tgbotapi.Message)
			}
			messages[update.Message.Chat.ID][strings.ToLower(update.Message.From.FirstName)] = update.Message
			messages[update.Message.Chat.ID][fmt.Sprintf("@%s", strings.ToLower(update.Message.From.UserName))] = update.Message
		}
	}
}

func randomCase(message string) string {
	var modMessage string
	message = strings.ToLower(message)
	for _, letter := range message {
		if weightedRandom() {
			modMessage += strings.ToUpper(string(letter))
		} else {
			modMessage += string(letter)
		}
	}
	return modMessage
}

func weightedRandom() bool {
	return prob[curRand.Intn(len(prob)-1)]
}
