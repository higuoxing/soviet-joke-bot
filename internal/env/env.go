package env

import (
	"errors"
	"os"
	"strconv"
)

// Env contains some essential environment parameters.
type Env struct {
	telegramBotToken string
	debugBot         bool
	botTimeout       int
	databasePath     string
}

// GetEnv returns environment parameters.
func GetEnv() (*Env, error) {
	e := &Env{}

	token := os.Getenv("TELEGRAM_API_TOKEN")
	if token == "" {
		return nil, errors.New("lacking of environment parameter TELEGRAM_API_TOKEN")
	}
	e.telegramBotToken = token

	if debug := os.Getenv("DEBUG_TELEGRAM_BOT"); debug == "true" {
		e.debugBot = true
	} else {
		e.debugBot = false
	}

	botTimeout := os.Getenv("BOT_TIMEOUT")
	if botTimeout != "" {
		timeout, err := strconv.Atoi(botTimeout)
		if err != nil {
			return nil, errors.New("environment parameter BOT_TIMEOUT should be integer")
		}

		e.botTimeout = timeout
	} else {
		e.botTimeout = 60
	}

	databasePath := os.Getenv("DB_PATH")
	if databasePath == "" {
		e.databasePath = "./data.db"
	} else {
		e.databasePath = databasePath
	}

	return e, nil
}

// TelegramBotToken returns telegram bot token.
func (e *Env) TelegramBotToken() string {
	return e.telegramBotToken
}

// DebugBot returns boolean that indicates whether to debug bot.
func (e *Env) DebugBot() bool {
	return e.debugBot
}

func (e *Env) DatabasePath() string {
	return e.databasePath
}
