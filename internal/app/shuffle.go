package app

import (
	"errors"
	"math/rand"
	"tgBot/internal/domain/entities"
	"time"
)

func (l *LobbyService) ShuffleLobby(n int, user *entities.User) ([][]*entities.User, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	code, ok := l.userToLobby[user.Name()]
	if !ok {
		return nil, errors.New("code not found")
	}
	lobby, exists := l.lobbies[code]
	if !exists {
		return nil, errors.New("lobby not found")
	}

	if len(lobby.Members) == 0 {
		return nil, errors.New("в лобби нет участников")
	}

	shuffles := make([][]*entities.User, 0, n)

	// используем копию чтобы не менять порядок оригинала
	original := make([]*entities.User, len(lobby.Members))
	copy(original, lobby.Members)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		perm := make([]*entities.User, len(original))
		copy(perm, original)
		r.Shuffle(len(perm), func(i, j int) {
			perm[i], perm[j] = perm[j], perm[i]
		})
		shuffles = append(shuffles, perm)
	}

	return shuffles, nil
}
