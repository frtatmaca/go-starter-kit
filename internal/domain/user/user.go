package domain

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
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
