package entities

type User struct {
	name   string
	chatId int64
	nick   string
}

func (u *User) ChatId() int64 {
	return u.chatId
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Nick() string {
	return u.nick
}

func NewUser(name string) *User {
	return &User{name: name}
}

func (u *User) WithChatId(chatId int64) *User {
	u.chatId = chatId
	return u
}

func (u *User) WithNick(nick string) *User {
	u.nick = nick
	return u
}
