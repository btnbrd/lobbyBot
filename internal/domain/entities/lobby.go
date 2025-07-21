package entities

import "errors"

type Lobby struct {
	Host    *User
	Code    string
	Members []*User
	Max     int
	Open    bool
}

func NewLobby(code string, max int) *Lobby {
	return &Lobby{
		Code: code,
		Max:  max,
	}
}

func (l *Lobby) WithHost(host *User) *Lobby {
	l.Host = host
	return l
}

func (l *Lobby) Closed() *Lobby {
	l.Open = false
	return l
}

func (l *Lobby) Opened() *Lobby {
	l.Open = true
	return l
}

func (l *Lobby) AddUser(u *User) error {
	if len(l.Members) >= l.Max {
		return errors.New("maximum number of members exceeded")
	}
	l.Members = append(l.Members, u)
	return nil
}
