package dto

import "learn/ginEssential/models"

type UserDto struct {
	Name string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user models.User) UserDto {
	return UserDto{
		Name: user.Name,
		Telephone: user.Telephone,
	}
}