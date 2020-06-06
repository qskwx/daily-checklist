package main

import (
	bot "daily-checklist/src/interface"
)

func main() {
	bot, _ := bot.BotFactory(telegramAPIKey)
	bot.Loop()
}
