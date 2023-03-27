package service

import (
	"NotePad/serializer"
)

func (userRegister *RegisterUserService) Register() serializer.Response {

	return serializer.Response{
		Status: 1,
	}
}
