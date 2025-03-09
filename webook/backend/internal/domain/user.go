package domain

type User struct {
	Email    string
	PassWord string
}

func NewUser(email, password string) User {
	return User{Email: email, PassWord: password}
}
