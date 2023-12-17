package serializer

import "todo_list/model"

type User struct {
	ID       uint   `json:"id" form:"id" example:"1"`
	UserName string `json:"username" form:"username" example:"NoName"`
	Status   string `json:"status" form:"status"`
	CreateAt int64  `json:"create_at" form:"create_at"`
}

func BuildUser(user model.User) User {

	return User{
		ID:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
	}
}
