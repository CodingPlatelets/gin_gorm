package Dto

import "github.com/WenkanHuang/gin_gorm/Model"

type UserDto struct {
	Name string `json:"name,omitempty"`
}

func ToUserDto(user Model.User) UserDto {
	return UserDto{
		Name: user.Name,
	}
}
