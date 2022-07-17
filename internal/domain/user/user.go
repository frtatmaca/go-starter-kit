package domain

type User struct {
	ID       int
	Email    string
	Password string
}

func NewUser(id int, email string, password string) *User {
	return &User{
		ID:       id,
		Email:    email,
		Password: password,
	}
}

func (u *User) GetEmail() string {
	return u.Email
}
