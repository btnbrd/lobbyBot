package app

import (
	"errors"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"tgBot/internal/domain/entities"
)

type LobbyService struct {
	mu          sync.Mutex
	lobbies     map[string]*entities.Lobby
	userToLobby map[string]string
	letters     []rune
}

func NewLobbyService() *LobbyService {
	return &LobbyService{
		lobbies:     make(map[string]*entities.Lobby),
		userToLobby: make(map[string]string),
		letters:     []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
	}
}

func (s *LobbyService) leaveNoLock(user *entities.User) error {
	code, ok := s.userToLobby[user.Name()]
	if !ok {
		return errors.New("ты не в лобби")
	}
	lobby, exists := s.lobbies[code]
	if !exists {
		delete(s.userToLobby, user.Name())
		return nil
	}

	newMembers := make([]*entities.User, 0, len(lobby.Members)-1)
	for _, id := range lobby.Members {
		if id.Name() != user.Name() {
			newMembers = append(newMembers, id)
		}
	}
	lobby.Members = newMembers

	delete(s.userToLobby, user.Name())

	if len(lobby.Members) == 0 {
		go func() {
			time.Sleep(30 * time.Second)
			delete(s.lobbies, code)
		}()
	}
	return nil
}

func (s *LobbyService) randCode() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 4)
	for i := range b {
		b[i] = s.letters[rand.Intn(len(s.letters))]
	}
	return string(b)
}

func formatUser(id int64) string {
	return "id" + strconv.FormatInt(id, 10)
}
