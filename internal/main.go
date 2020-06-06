package main

import (
	bot "daily-checklist/src/interface"
)

const telegramAPIKey string = "875847790:AAEKNbIJ-U8X99I1JTpgAuDF15ZSqy9m8DI"

func main() {
	bot, _ := bot.BotFactory(telegramAPIKey)
	bot.Loop()
}
