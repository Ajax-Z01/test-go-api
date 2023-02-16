package response

import (
	"encoding/json"
	"test-go-api/entity"
)

type User struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func MapResponseListUser(listUser []*entity.User) ([]byte, error) {
	var dataJson []*User

	for _, user := range listUser {
		var jsonSingle = &User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
		}

		dataJson = append(dataJson, jsonSingle)
	}

	byteJson, err := json.Marshal(dataJson)

	if err != nil {
		return nil, err
	}

	return byteJson, nil
}
