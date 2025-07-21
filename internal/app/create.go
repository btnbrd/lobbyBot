package app

import "tgBot/internal/domain/entities"

const maxCount = 10

func (s *LobbyService) CreateLobby(user *entities.User) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.leaveNoLock(user)

	var code string
	for {
		code = s.randCode()
		if _, exists := s.lobbies[code]; !exists {
			break
		}
	}

	lobby := entities.NewLobby(code, 10).WithHost(user).Closed()
	lobby.AddUser(user)

	s.lobbies[code] = lobby
	s.userToLobby[user.Name()] = code

	return code
}
