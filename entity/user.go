package entity

type User struct {
	ID       int    `json:"id,omitempty"`
	USERNAME string `json:"username,omitempty"`
	EMAIL    string `json:"email,omitempty"`
}

type UserDTO struct {
	ID       int
	Username string
	Email    string
}

func CreateUser(dto UserDTO) *User {
	return &User{
		ID:       dto.ID,
		USERNAME: dto.Username,
		EMAIL:    dto.Email,
	}
}

func UpdateUser(user *User, dto UserDTO) {
	user.ID = dto.ID
	user.USERNAME = dto.Username
	user.EMAIL = dto.Email
}

func (b *User) GetID() int {
	return b.ID
}

func (b *User) GetUsername() string {
	return b.USERNAME
}

func (b *User) GetEmail() string {
	return b.EMAIL
}
