package dto

import "github.com/WenkanHuang/gin_gorm/model"

type UserDto struct {
	Name string `json:"name,omitempty"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name: user.Name,
	}
}
