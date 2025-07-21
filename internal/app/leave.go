package app

import "tgBot/internal/domain/entities"

func (s *LobbyService) LeaveLobby(user *entities.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.leaveNoLock(user)
}
