package app

import (
	"errors"
	"tgBot/internal/domain/entities"
)

func (s *LobbyService) JoinRandom(user *entities.User) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.leaveNoLock(user)

	for code, lobby := range s.lobbies {
		if lobby.Open {
			if err := lobby.AddUser(user); err == nil {
				s.userToLobby[user.Name()] = code
				return code, nil
			}
		}
	}
	return "", errors.New("нет доступных лобби")
}
