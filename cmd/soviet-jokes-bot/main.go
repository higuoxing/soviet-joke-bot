package main

import (
	"fmt"
	"log"

	"github.com/vgxbj/soviet-jokes-bot/internal/bot"
	"github.com/vgxbj/soviet-jokes-bot/internal/db"
	"github.com/vgxbj/soviet-jokes-bot/internal/env"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	e, err := env.GetEnv()
	if err != nil {
		log.Panic(err)
	}

	d, err := db.InitDatabase(e.DatabasePath())
	if err != nil {
		log.Panic(err)
	}

	bot, err := bot.NewBot(e)
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we should leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "random":
			joke, err := d.Random()
			if err != nil {
				msg.Text = fmt.Sprintf("%v", err)
			} else {
				msg.Text = joke
			}
		default:
			msg.Text = "我不明白"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
