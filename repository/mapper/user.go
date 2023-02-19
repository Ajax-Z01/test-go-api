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

func MapToDTO(listEntity []*entity.User) []entity.UserDTO {
	var listDTO []entity.UserDTO
	for _, user := range listEntity {
		dto := entity.UserDTO{
			ID:       user.ID,
			Username: user.USERNAME,
			Email:    user.EMAIL,
		}
		listDTO = append(listDTO, dto)
	}
	return listDTO
}

func MapToEntitySingle(dto entity.UserDTO) *entity.User {
	return entity.CreateUser(dto)
}

func MapToDTOSingle(user *entity.User) entity.UserDTO {
	return entity.UserDTO{
		ID:       user.ID,
		Username: user.USERNAME,
		Email:    user.EMAIL,
	}
}

func MapToEntityUpdate(user *entity.User, dto entity.UserDTO) {
	entity.UpdateUser(user, dto)
}
