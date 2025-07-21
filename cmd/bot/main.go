package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load() // локально читает .env, Railway это игнорирует.

	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_TOKEN не найден")
	}

	log.Println("TOKEN успешно загружен:", token[:5]) // выводим часть токена для проверки
}
