package main

import (
	"log"
	"os"
	"tgBot/internal/adapters/telegram"
	"tgBot/internal/app"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_TOKEN не найден")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Запущен бот: %s", bot.Self.UserName)

	service := app.NewLobbyService()
	handler := telegram.NewHandler(bot, service)
	router := telegram.NewRouter(bot, handler)
	router.StartPolling()
}
