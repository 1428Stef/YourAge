package main

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("TOKEN")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "start":
			msg.Text = "Hi, please enter your date of birth in the format YYYY-MM-DD"
		case "help":
			msg.Text = "Just enter your date of birth 2006-01-02"
		default:
			if len(update.Message.Text) == 10 {
				age, err := calculateAge(update.Message.Text)
				if err != nil {
					msg.Text = "Invalid date format"
				} else {
					msg.Text = fmt.Sprintf("You are %d years old", age)
				}
			} else {
				msg.Text = "Please provide your correct date of birth in YYYY-MM-DD format."
			}
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

func calculateAge(yourbirth string) (int, error) {
	now := time.Now()

	your, err := time.Parse("2006-01-02", yourbirth)
	if err != nil {
		return 0, err
	}

	age := now.Year() - your.Year()
	if now.YearDay() < your.YearDay() {
		age--
	}

	return age, nil
}
