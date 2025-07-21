package app

import "tgBot/internal/domain/entities"

func (s *LobbyService) Register(user *entities.User) {
	s.userToChat[user.Name()] = user.ChatId()
}
