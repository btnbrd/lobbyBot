package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgBot/internal/app"
	"tgBot/internal/domain/entities"
)

func Whoami(service *app.LobbyService, user *entities.User, msg *tgbotapi.MessageConfig) {
	code, members, err := service.GetUserLobby(user)
	if err != nil {
		msg.Text = "Ошибка: " + err.Error()
	} else {
		msg.Text = "Ты в лобби " + code + ". Участники: " + members
	}
}
