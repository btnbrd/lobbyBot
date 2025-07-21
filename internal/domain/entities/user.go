package entities

type User struct {
	// id   int64
	name string
}

// func (u User) Id() string {
// 	return u.name
// }

func (u User) Name() string {
	return u.name
}

func NewUser(name string) *User {
	return &User{name}
}
