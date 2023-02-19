package entity

type User struct {
	id       int
	username string
	email    string
	password string
}

type UserDTO struct {
	ID       int
	Username string
	Email    string
	Password string
}

func CreateUser(dto UserDTO) *User {
	user := &User{
		id:       dto.ID,
		username: dto.Username,
		email:    dto.Email,
		password: dto.Password,
	}

	return user
}

func (b *User) GetID() int {
	return b.id
}

func (b *User) GetUsername() string {
	return b.username
}

func (b *User) GetEmail() string {
	return b.email
}
