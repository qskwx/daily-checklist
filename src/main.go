package main

import (
	bot "daily-checklist/src/interface"
)

const telegramAPIKey string = ""

func main() {
	bot, _ := bot.BotFactory(telegramAPIKey)
	bot.Loop()
}
