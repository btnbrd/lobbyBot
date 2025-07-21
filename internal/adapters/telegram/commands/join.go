package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"tgBot/internal/app"
	"tgBot/internal/domain/entities"
)

func Join(service *app.LobbyService, user *entities.User, code string, msg *tgbotapi.MessageConfig) {
	err := service.JoinLobby(code, user)
	if err != nil {
		msg.Text = "Ошибка: " + err.Error()
	} else {
		msg.Text = "Успешно присоединился к лобби: " + strings.ToUpper(code)
	}
}
