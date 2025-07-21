package app

import (
	"errors"
	"strings"
	"tgBot/internal/domain/entities"
)

func (s *LobbyService) ConnectUser(host *entities.User, user *entities.User) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	code, ok := s.userToLobby[host.Name()]
	if !ok {
		return "", errors.New("<UNK>")
	}
	code = strings.ToUpper(code)
	lobby := s.lobbies[code]

	if lobby.Host.Name() != host.Name() {
		return "", errors.New("You are not the host")
	}

	s.leaveNoLock(user)
	if err := lobby.AddUser(user); err != nil {
		return "", err
	}
	s.userToLobby[user.Name()] = code

	return code, nil
}
