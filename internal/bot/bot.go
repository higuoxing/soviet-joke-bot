package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/vgxbj/soviet-jokes-bot/internal/env"
)

// Bot is alias for *tgbotapi.BotAPI
type Bot = tgbotapi.BotAPI

// NewBot returns a bot.
func NewBot(e *env.Env) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(e.TelegramBotToken())
	if err != nil {
		return nil, err
	}

	// Set debug arguments.
	bot.Debug = e.DebugBot()

	if e.DebugBot() {
		log.Printf("Authorized on account %s", bot.Self.UserName)
	}

	return bot, nil
}
