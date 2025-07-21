package commands

import (
	"fmt"
	"strconv"
	"strings"
	"tgBot/internal/app"
	"tgBot/internal/domain/entities"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Shuffle(service *app.LobbyService, user *entities.User, arg string, msg *tgbotapi.MessageConfig) {
	n, err := strconv.Atoi(strings.TrimSpace(arg))
	if err != nil || n <= 0 {
		msg.Text = "Пожалуйста, укажи положительное число. Пример: /shuffle 3"
		return
	}

	permutations, err := service.ShuffleLobby(n, user)
	if err != nil {
		msg.Text = "Ошибка: " + err.Error()
		return
	}
	text := ""
	for i, perm := range permutations {
		text += fmt.Sprintf("%d: ", i+1)
		for _, u := range perm {
			text += u.Name() + " "
		}
		text += "\n"
	}
	msg.Text = text
}
