package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var prob = []bool{true, true, false, false, false}

var curRand *rand.Rand

func init() {
	curRand = rand.New(rand.NewSource(time.Now().Unix()))
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

	names := getNames()

	if len(names) < 1 {
		log.Fatal("no names in env")
	}

	for update := range updates {
		if update.Message == nil || update.Message.Text == "" || !stringInSlice(update.Message.From.UserName, names) || !weightedRandom() { // ignore any non-Message Updates
			continue
		}

		gif := tgbotapi.NewAnimationShare(update.Message.Chat.ID, "CgACAgQAAxkBAAMdYGDU-aGj-BAjTa0dGviX35Z0gP4AAjYCAALYL5VSHxQhj1TU2cYeBA")
		gif.Caption = randomCase(update.Message.Text)
		gif.ReplyToMessageID = update.Message.MessageID
		_, err = bot.Send(gif)

		if err != nil {
			log.Println(err)
		}
	}
}

func getNames() []string {
	nameEnv := os.Getenv("TELEGRAM_TARGETS")
	return strings.Split(nameEnv, ",")
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

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}