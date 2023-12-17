package service

import (
	"fmt"
	"todo_list/model"
	"todo_list/serializer"
	utils "todo_list/token"

	"github.com/jinzhu/gorm"
)

type UserService struct {
	UserName string `form:"username" json:"username"`
	PassWord string `form:"password" json:"password"`
}

func (service *UserService) Register() serializer.Response {
	fmt.Println("Register running")
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count) //连接数据库并查找是否已经注册
	if count == 1 {

		return serializer.Response{
			Message: "您已注册，请勿重复注册",
			Status:  400,
		}
	}
	user.UserName = service.UserName
	//密码加密
	err := user.Encrypt_password(service.PassWord)
	if err != nil {
		return serializer.Response{
			Message: err.Error(),
			Status:  400,
		}
	}
	if err2 := model.DB.Create(&user).Error; err2 != nil {
		return serializer.Response{
			Message: "数据库操作错误",
			Status:  500,
		}
	}
	return serializer.Response{
		Message: "注册成功",
		Status:  200,
	}
}
func (service *UserService) Login() serializer.Response {
	var user model.User
	//数据库中查找用户
	if err := model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status:  400,
				Message: "查无此人，请先注册",
			}
		}

		//查找不到但并非用户不存在
		return serializer.Response{
			Status:  500,
			Message: "数据库错误",
		}
	}

	if !user.Compare_password(service.PassWord) {
		return serializer.Response{
			Status:  400,
			Message: "密码错误",
		}
	}
	//发token，为了需要其他功能需要身份验证的功能给前端存储的
	//例如创建备忘录就需要token，否则不知道谁是备忘录创建者
	token, err := utils.GenerateToken(user.ID, service.UserName, service.PassWord)
	if err != nil {
		return serializer.Response{
			Status:  500,
			Message: "token签发错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
		Message: "登录成功",
	}

}
