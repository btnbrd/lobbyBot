package main

import (
	"log"
	"os"
)

func main() {
	for _, e := range os.Environ() {
		log.Println("ENV:", e)
	}

	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_TOKEN не найден")
	}
}
