package mapper

import "test-go-api/entity"

func MapToEntity(listDTO []entity.UserDTO) []*entity.User {
	var listUser []*entity.User
	for _, dto := range listDTO {
		user := entity.CreateUser(dto)
		listUser = append(listUser, user)
	}

	return listUser
}
