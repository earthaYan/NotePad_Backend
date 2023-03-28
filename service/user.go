package service

import (
	"NotePad/serializer"
	"fmt"
)

func (userRegister *RegisterUserService) Register() serializer.Response {
	fmt.Println(userRegister)
	return serializer.Response{
		Status: 200,
		Msg: "用户注册成功",
	}
}
