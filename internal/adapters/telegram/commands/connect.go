package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"tgBot/internal/app"
	"tgBot/internal/domain/entities"
)

func Connect(service *app.LobbyService, user *entities.User, args string, msg *tgbotapi.MessageConfig) {
	arg := strings.TrimSpace(args)
	sub := entities.NewUser(arg)
	code, err := service.ConnectUser(user, sub)
	if err != nil {
		msg.Text = "Ошибка: " + err.Error()
	} else {
		msg.Text = sub.Name() + " успешно присоединился к лобби: " + strings.ToUpper(code)
	}
}
