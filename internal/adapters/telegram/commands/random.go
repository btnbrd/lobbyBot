package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgBot/internal/app"
	"tgBot/internal/domain/entities"
)

func Random(service *app.LobbyService, user *entities.User, msg *tgbotapi.MessageConfig) {
	code, err := service.JoinRandom(user)
	if err != nil {
		msg.Text = "Ошибка: " + err.Error()
	} else {
		msg.Text = "Ты присоединился к случайному лобби: " + code
	}
}
