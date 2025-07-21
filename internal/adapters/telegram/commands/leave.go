package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgBot/internal/app"
	"tgBot/internal/domain/entities"
)

func Leave(service *app.LobbyService, user *entities.User, msg *tgbotapi.MessageConfig) {
	err := service.LeaveLobby(user)
	if err != nil {
		msg.Text = "Ошибка: " + err.Error()
	} else {
		msg.Text = "Ты вышел из лобби"
	}
}
