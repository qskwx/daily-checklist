package main

import (
	bot "daily-checklist/internal/interface"
)

func main() {
	bot, _ := bot.BotFactory(telegramAPIKey)
	bot.Loop()
}
