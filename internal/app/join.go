package app

import (
	"errors"
	"strings"
	"tgBot/internal/domain/entities"
)

func (s *LobbyService) JoinLobby(code string, user *entities.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	code = strings.ToUpper(code)
	lobby, ok := s.lobbies[code]
	if !ok {
		return errors.New("лобби не найдено")
	}
	if len(lobby.Members) >= lobby.Max {
		return errors.New("лобби заполнено")
	}

	s.leaveNoLock(user)

	lobby.AddUser(user)
	s.userToLobby[user.Name()] = code
	return nil
}
