package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Router struct {
	bot     *tgbotapi.BotAPI
	handler *Handler
}

func NewRouter(bot *tgbotapi.BotAPI, handler *Handler) *Router {
	return &Router{
		bot:     bot,
		handler: handler,
	}
}

func (r *Router) StartPolling() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := r.bot.GetUpdatesChan(u)

	for update := range updates {
		go r.handler.Handle(update)
	}
}
