package iface

import (
	"daily-checklist/internal/session"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot      *tgbotapi.BotAPI
	sessions session.Sessions
}

func BotFactory(APIKey string) (Bot, error) {
	bot, err := tgbotapi.NewBotAPI(APIKey)

	if err != nil {
		log.Panic(err)
		return Bot{}, err
	}

	// bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	sessions := session.SessionsFabric()
	return Bot{bot: bot, sessions: sessions}, nil
}

func (bot *Bot) Loop() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
		return
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		if update.Message.From.UserName != "vo0xr0c" { // ignore all exclude me
			log.Printf("Request form [%s] ignored", update.Message.From.UserName)
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		reply := bot.processUpdate(update)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		bot.bot.Send(msg)
	}
}

func (bot *Bot) processUpdate(update tgbotapi.Update) (reply string) {
	ss, err := bot.sessions.Session(update.Message.From.UserName, time.Now())
	if err != nil {
		return err.Error()
	}

	switch update.Message.Text {
	case "?":
		reply = constructMessage(ss)
	default:
		if resolve := ss.SetDone(update.Message.Text); resolve != nil {
			reply = resolve.Error()
		} else {
			reply = constructMessage(ss)
		}
	}
	return
}
