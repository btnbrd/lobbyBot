package ports

import (
	lobby "tgBot/internal/domain/entities"
)

type LobbyRepository interface {
	Create(lobby *lobby.Lobby) error
	Join(code string, userID int64) error
	Leave(userID int64) error
	GetByUser(userID int64) (*lobby.Lobby, error)
	FindAvailable() (*lobby.Lobby, error)
}
