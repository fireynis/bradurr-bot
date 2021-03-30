package main

import (
	"fmt"
	"go.uber.org/zap"
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

	if os.Getenv("BOT_ENV") != "production" {
		bot.Debug = true
	}

	starter, err := zap.NewProduction()

	if err != nil {
		log.Panic(err)
	}

	logger := starter.Sugar()

	logger.Infow("Authorized as user", zap.String("bot_username", bot.Self.UserName))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			if strings.ToLower(update.Message.Command()) == "getgif" && update.Message.Chat.IsPrivate() {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				if message, ok := messages[update.Message.Chat.ID][strings.ToLower(update.Message.From.FirstName)]; !ok || message.Document == nil || message.Document.FileID == "" {
					msg.Text = "We ain't found shit"
				} else {
					msg.Text = message.Document.FileID
					_, _ = bot.Send(msg)
				}
			}
			var gif tgbotapi.AnimationConfig
			switch strings.ToLower(update.Message.Command()) {
			case "durr":
				gif = tgbotapi.NewAnimationShare(update.Message.Chat.ID, "CgACAgQAAxkBAAMdYGDU-aGj-BAjTa0dGviX35Z0gP4AAjYCAALYL5VSHxQhj1TU2cYeBA")
				if message, ok := messages[update.Message.Chat.ID][strings.ToLower(update.Message.CommandArguments())]; !ok {
					gif.Caption = randomCase(fmt.Sprintf("hurr durr %s, %s has no messages in here yet", update.Message.From.FirstName, strings.ToTitle(update.Message.CommandArguments())))
				} else {
					gif.Caption = randomCase(message.Text)
					gif.ReplyToMessageID = message.MessageID
				}
			case "cowboy":
				gif = tgbotapi.NewAnimationShare(update.Message.Chat.ID, "AAMCBAADGQEAAgRsYGJwCA_1mY1qydjoehoFzS0S4ZwAAnACAALAMa1SBU-G3K6fQJXYIpQmXQADAQAHbQADb18AAh4E")
				if message, ok := messages[update.Message.Chat.ID][strings.ToLower(update.Message.CommandArguments())]; !ok {
					gif.Caption = randomCase(fmt.Sprintf("Nobody here, but still horny %s", update.Message.From.FirstName))
				} else {
					gif.Caption = message.Text
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
					gif.Caption = fmt.Sprintf("%s, you thirsty?", update.Message.From.FirstName)
				} else {
					gif.ReplyToMessageID = message.MessageID
				}
			case "danger":
				gif = tgbotapi.NewAnimationShare(update.Message.Chat.ID, "CgACAgQAAxkBAAIEsGBie58bKOeSXEHrVrvh-Gta5Wk_AAIrAgACiX-VUpZXPFLOB0VfHgQ")
				if message, ok := messages[update.Message.Chat.ID][strings.ToLower(update.Message.CommandArguments())]; !ok {
					gif.Caption = fmt.Sprintf("%s, YOU DON'T REALLY EXIST", update.Message.From.FirstName)
				} else {
					gif.ReplyToMessageID = message.MessageID
				}
			case "nerd":
				gif = tgbotapi.NewAnimationShare(update.Message.Chat.ID, "CgACAgQAAxkBAAIEs2BifFqTRTFAL-cM4Bhv66NW-FBoAAIaAgACP7CcUooP90RkGHy0HgQ")
				if message, ok := messages[update.Message.Chat.ID][strings.ToLower(update.Message.CommandArguments())]; !ok {
					gif.Caption = fmt.Sprintf("%s, no user name? Fuckin nerd.", update.Message.From.FirstName)
				} else {
					gif.ReplyToMessageID = message.MessageID
				}
			case "yousuck":
				gif = tgbotapi.NewAnimationShare(update.Message.Chat.ID, "CgACAgQAAxkBAAIEv2BifdkT3G_j-GANz0CufXCvdHZuAAITAgACuhqsUcCIL8xUN9IfHgQ")
				if message, ok := messages[update.Message.Chat.ID][strings.ToLower(update.Message.CommandArguments())]; !ok {
					gif.Caption = fmt.Sprintf("%s, you suck donkey dick", update.Message.From.FirstName)
				} else {
					gif.ReplyToMessageID = message.MessageID
				}
			case "donkey":
				gif = tgbotapi.NewAnimationShare(update.Message.Chat.ID, "CgACAgQAAxkBAAIEwmBifh6-60ypp0HnpFvbfLefd54TAAJHAgACv26lUrwh-zTY0RjfHgQ")
				if message, ok := messages[update.Message.Chat.ID][strings.ToLower(update.Message.CommandArguments())]; !ok {
					gif.Caption = fmt.Sprintf("%s, you're the donkey", update.Message.From.FirstName)
				} else {
					gif.ReplyToMessageID = message.MessageID
				}
			case "drool":
				gif = tgbotapi.NewAnimationShare(update.Message.Chat.ID, "CgACAgQAAxkBAAIEi2BidvEg3yU13GAk3NMJV_o0PirQAAIkAgAChh-UUt_G8LPDSbVVHgQ")
				if message, ok := messages[update.Message.Chat.ID][strings.ToLower(update.Message.CommandArguments())]; !ok {
					gif.Caption = fmt.Sprintf("Drooling on yourself %s?", update.Message.From.FirstName)
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
			case "stfu":
				gif = tgbotapi.NewAnimationShare(update.Message.Chat.ID, "CgACAgQAAxkBAAIEt2BifSIpPjUgjgABIlXtU6dV7fsOeQACUAIAAp49jFJFzO_VtudJUh4E")
				if message, ok := messages[update.Message.Chat.ID][strings.ToLower(update.Message.CommandArguments())]; !ok {
					gif.Caption = fmt.Sprintf("No one cares, STFU %s!", update.Message.From.FirstName)
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
