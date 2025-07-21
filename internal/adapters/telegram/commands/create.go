package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgBot/internal/app"
	"tgBot/internal/domain/entities"
)

func Create(service *app.LobbyService, user *entities.User, msg *tgbotapi.MessageConfig) {
	code := service.CreateLobby(user)
	msg.Text = "Создано лобби с кодом: " + code
}
