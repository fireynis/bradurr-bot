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

var logAll bool

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
		logAll = true
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
		if logAll {
			if update.Message.ReplyToMessage != nil {
				logger.Infow(
					"reply to message occurred",
					zap.Int("message_id", update.Message.ReplyToMessage.MessageID),
				)
			}
		}
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			commandName := strings.ToLower(update.Message.Command())
			var gif tgbotapi.AnimationConfig
			switch commandName {
			case "getgif":
				if !update.Message.Chat.IsPrivate() {
					continue
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				replyMessage, err := getMessage(&update)
				if err != nil || replyMessage.Document == nil || replyMessage.Document.FileID == "" {
					msg.Text = "We ain't found shit"
				} else {
					msg.Text = replyMessage.Document.FileID
					_, _ = bot.Send(msg)
				}
				continue
			case "durr":
				gif = gifGenerate(
					&update,
					"CgACAgQAAxkBAAIF9WBjq7WaXMQmzKv6dNMlyuTXJM19AAI2AgAC2C-VUh8UIY9U1NnGHgQ",
					fmt.Sprintf(
						"hurr durr %s, %s has no messages in here yet",
						update.Message.From.FirstName,
						strings.ToTitle(update.Message.CommandArguments()),
					),
					true,
					true,
				)
				break
			case "cowboy":
				gif = gifGenerate(
					&update,
					"CgACAgQAAxkBAAIHJmBrKSW6-BfQEj0y7yw9eOmVrI9YAAJwAgACwDGtUgVPhtyun0CVHgQ",
					fmt.Sprintf("Nobody here, but still horny %s", update.Message.From.FirstName),
					false,
					false,
				)
				break
			case "fuckyou":
				gif = gifGenerate(
					&update,
					"CgACAgQAAxkBAAIF72Bjq4Am72Oc5PYPAfAd_Py-loqNAAJxAgACeqeMUjwXqK1QN7qGHgQ",
					fmt.Sprintf("Hey, hey %s. Fuck you, no messages from %s yet", update.Message.From.FirstName, strings.ToTitle(update.Message.CommandArguments())),
					false,
					false,
				)
				break
			case "sploosh":
				gif = gifGenerate(
					&update,
					"CgACAgQAAxkBAAIBVWBhLj-iPzQzDDiwNi75UT7bTKk2AAKDAgACg0mNUnjNm9ca7IUmHgQ",
					fmt.Sprintf("You splooshed for nothing %s, %s has no messages in here yet", update.Message.From.FirstName, strings.ToTitle(update.Message.CommandArguments())),
					false,
					false,
				)
				break
			case "goldstar":
				gif = gifGenerate(
					&update,
					"CgACAgQAAxkBAAICOGBiJptsJgqu7CjI7wglyFpry5rFAAJUAgACu-uVUt2pTnE70uewHgQ",
					fmt.Sprintf("%s gets a dumbass gold star for no name sent, what a dumbass!", update.Message.From.FirstName),
					false,
					false,
				)
				break
			case "magma":
				gif = gifGenerate(
					&update,
					"CgACAgEAAxkBAAIEcmBicgTC9R0ZMNginJB6ewKudpfhAAJsAQACXyUQR0jwan3HzlslHgQ",
					fmt.Sprintf("Not fucking lava dude %s", update.Message.From.FirstName),
					false,
					false,
				)
				break
			case "boobs":
				gif = gifGenerate(
					&update,
					"CgACAgQAAxkBAAIEpmBieuj2CmF42NGZFMj5-LViv7sNAAIqAgACno3sUSMPJOU4zQYYHgQ",
					fmt.Sprintf("%s, you thirsty?", update.Message.From.FirstName),
					false,
					false,
				)
				break
			case "danger":
				gif = gifGenerate(
					&update,
					"CgACAgQAAxkBAAIEsGBie58bKOeSXEHrVrvh-Gta5Wk_AAIrAgACiX-VUpZXPFLOB0VfHgQ",
					fmt.Sprintf("%s, eat to much paste?", update.Message.From.FirstName),
					false,
					false,
				)
				break
			case "nerd":
				gif = gifGenerate(
					&update,
					"CgACAgQAAxkBAAIEs2BifFqTRTFAL-cM4Bhv66NW-FBoAAIaAgACP7CcUooP90RkGHy0HgQ",
					fmt.Sprintf("%s, no user name? Fuckin nerd.", update.Message.From.FirstName),
					false,
					false,
				)
				break
			case "yousuck":
				gif = gifGenerate(
					&update,
					"CgACAgQAAxkBAAIEv2BifdkT3G_j-GANz0CufXCvdHZuAAITAgACuhqsUcCIL8xUN9IfHgQ",
					fmt.Sprintf("%s, you suck donkey dick", update.Message.From.FirstName),
					false,
					false,
				)
				break
			case "donkey":
				gif = gifGenerate(
					&update,
					"CgACAgQAAxkBAAIEwmBifh6-60ypp0HnpFvbfLefd54TAAJHAgACv26lUrwh-zTY0RjfHgQ",
					fmt.Sprintf("%s, you're the donkey", update.Message.From.FirstName),
					false,
					false,
				)
				break
			case "drool":
				gif = gifGenerate(
					&update,
					"CgACAgQAAxkBAAIEi2BidvEg3yU13GAk3NMJV_o0PirQAAIkAgAChh-UUt_G8LPDSbVVHgQ",
					fmt.Sprintf("Drooling on yourself %s?", update.Message.From.FirstName),
					false,
					false,
				)
				break
			case "wtf":
				gif = gifGenerate(
					&update,
					"CgACAgQAAxkBAAICbWBiMQW6JtqhSaQRUI3s_EXSI1TuAAJJAgACiAOUUrtigLpU3wfdHgQ",
					fmt.Sprintf("Not event an actual username WTF %s", update.Message.From.FirstName),
					false,
					false,
				)
				break
			case "stfu":
				gif = gifGenerate(
					&update,
					"CgACAgQAAxkBAAIEt2BifSIpPjUgjgABIlXtU6dV7fsOeQACUAIAAp49jFJFzO_VtudJUh4E",
					fmt.Sprintf("No one cares, STFU %s!", update.Message.From.FirstName),
					false,
					false,
				)
				break
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

func gifGenerate(update *tgbotapi.Update, fileId, insult string, randomResponseCase, replyWithText bool) tgbotapi.AnimationConfig {
	var gif tgbotapi.AnimationConfig
	gif = tgbotapi.NewAnimationShare(update.Message.Chat.ID, fileId)
	replyMessage, err := getMessage(update)
	if err != nil {
		if randomResponseCase {
			gif.Caption = randomCase(insult)
		} else {
			gif.Caption = insult
		}
		return gif
	}
	if replyWithText {
		if randomResponseCase {
			gif.Caption = randomCase(replyMessage.Text)
		} else {
			gif.Caption = replyMessage.Text
		}
	}
	gif.ReplyToMessageID = replyMessage.MessageID
	return gif
}

func getMessage(update *tgbotapi.Update) (*tgbotapi.Message, error) {
	if update.Message.ReplyToMessage != nil {
		return update.Message.ReplyToMessage, nil
	}
	message, ok := messages[update.Message.Chat.ID][strings.ToLower(update.Message.CommandArguments())]
	if !ok {
		return nil, fmt.Errorf("no message from user found")
	}
	return message, nil
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
