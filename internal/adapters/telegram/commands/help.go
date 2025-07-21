package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgBot/internal/domain/entities"
)

func Help(user *entities.User, msg *tgbotapi.MessageConfig) {
	msg.Text = "Привет! Я бот для создания лобби.\nДоступные команды:\n" +
		"/create — создать лобби\n" +
		"/join — присоединиться к лобби по коду (введи /join и затем код в следующем сообщении)\n" +
		"/join <код> — присоединиться сразу\n" +
		"/random — присоединиться к случайному лобби\n" +
		"/leave — покинуть лобби\n" +
		"/whoami — узнать текущее лобби и участников\n" +
		"/shuffle <n> — получить n перестановок текущего лобби"
}
