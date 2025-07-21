package app

import (
	"errors"
	"fmt"
	"strings"
	"tgBot/internal/domain/entities"
)

func (s *LobbyService) GetUserLobby(user *entities.User) (string, string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	code, ok := s.userToLobby[user.Name()]
	if !ok {
		return "", "", errors.New("ты не в лобби")
	}
	lobby, exists := s.lobbies[code]
	if !exists {
		return "", "", errors.New("лобби не найдено")
	}
	host := lobby.Host.Name()
	var members []string
	for _, user := range lobby.Members {
		members = append(members, user.Name())
	}
	return code, fmt.Sprintf("Host: %s, Members: %s", host, strings.Join(members, ", ")), nil
}
