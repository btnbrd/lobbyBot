package commands

import (
	"tgBot/internal/app"
	"tgBot/internal/domain/entities"
)

func Start(service *app.LobbyService, user *entities.User) {
	service.Register(user)
}
