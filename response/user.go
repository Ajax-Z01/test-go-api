package response

import (
	"encoding/json"
	"test-go-api/entity"
	_interface "test-go-api/entity/interface"
)

type JSONListUser struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func MapResponseListUser(listUser []*entity.User) ([]byte, error) {
	var dataJson []*JSONListUser
	for _, user := range listUser {
		var jsonSingle = EntityToJson(user)

		dataJson = append(dataJson, jsonSingle)
	}

	byteJson, err := json.Marshal(dataJson)
	if err != nil {
		return nil, err
	}
	return byteJson, nil
}

func EntityToJson(user _interface.UserInterface) *JSONListUser {
	jsonSingle := &JSONListUser{
		ID:       user.GetID(),
		Username: user.GetUsername(),
		Email:    user.GetEmail(),
	}
	return jsonSingle
}

func MapResponseUser(user *entity.User) ([]byte, error) {
	data := map[string]interface{}{
		"username": user.USERNAME,
		"email":    user.EMAIL,
	}
	return json.Marshal(data)
}
