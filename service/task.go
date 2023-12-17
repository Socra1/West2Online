package service

import (
	"fmt"
	"time"
	"todo_list/model"
	"todo_list/serializer"
)

type Create_UpdateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0表示未做

}

type SearchtaskService struct {
	Info     string `josn:"info" form:"info"`
	PageNum  int    `json:"page_num" form:"page_num"`
	PageSize int    `json:"page_size" form:"page_size"`
} //可以分页的查询结构体
type Un_or_FinishiATaskService struct {
	Status int `json:"status" form:"status"` //0表示未做

} //分页功能的实现的结构体
type Un_or_FinishiAllTaskService struct {
	Status int `json:"status" form:"status"` //0表示未做
}

type EmptyTask struct {
}

func (service *Create_UpdateTaskService) Create(id uint) serializer.Response {
	var user model.User
	code := 200
	model.DB.First(&user, id)
	fmt.Printf("service: %v\n", service)
	task := model.Task{

		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Status:    0,
		Content:   service.Content,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status:  code,
			Message: "创建备忘录失败",
		}
	}
	return serializer.Response{

		Status:  code,
		Message: "创建成功",
	}

} //创建一条备忘录

func (service *SearchtaskService) Search(uid uint) serializer.Response {
	var tasks []model.Task
	cnt := 0
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	//预加载到用户后在用户的基础上搜索所需信息再进行计数->分页->赋值
	err := model.DB.Model(&model.Task{}).
		Preload("User").
		Where("uid=?", uid).
		Where("title LIKE ? OR content LIKE ?", "%"+service.Info+"%", "%"+service.Info+"%"). //在内容与标题中模糊查询
		Count(&cnt).
		Limit(service.PageSize).Offset((service.PageNum - 1) * (service.PageSize)).
		Find(&tasks).Error

	if err == nil {
		return serializer.Response{
			Status:  200,
			Data:    serializer.BuildTasks(tasks),
			Message: "查询成功",
		}
	} else {
		return serializer.Response{
			Status:  400,
			Message: "查询失败",
		}
	}

} //查询操作
func (service *EmptyTask) Delete(tid string) serializer.Response {
	var task model.Task
	err := model.DB.Delete(&task, tid).Error
	if err == nil {
		return serializer.Response{
			Status:  200,
			Message: "删除成功",
		}
	} else {
		return serializer.Response{
			Status:  500,
			Message: "删除失败",
		}
	}
} //删除操作
func (service *Un_or_FinishiATaskService) Un_or_FinishOne(tid string) serializer.Response {
	var task model.Task
	err := model.DB.First(&task, tid).Error
	if err != nil {
		return serializer.Response{
			Status:  400,
			Message: "设置失败",
		}

	} else {

		task.Status = service.Status
		model.DB.Save(&task)
		return serializer.Response{
			Status:  200,
			Data:    serializer.BuildTask(task),
			Message: "设置成功",
		}
	}

}
func (service *Un_or_FinishiAllTaskService) Un_or_FinishAll(uid uint) serializer.Response {
	var tasks []model.Task
	cnt := 0
	err := model.DB.Model(&model.Task{}).
		Preload("User").Where("uid=?", uid).
		Count(&cnt).
		Find(&tasks).Error
	if err != nil {
		return serializer.Response{
			Status:  400,
			Message: "设置所有待办事务失败",
		}
	} else {
		for _, v := range tasks {
			v.Status = service.Status
			model.DB.Save(&v)
		}

		return serializer.Response{
			Status:  200,
			Data:    serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(cnt)),
			Message: "设置所有待办事务成功",
		}
	}

}
func (service *SearchtaskService) Show_Finished(uid uint) serializer.Response {
	var tasks []model.Task
	cnt := 0
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	//预加载到用户后在用户的基础上搜索所需信息再进行计数->分页->赋值
	err := model.DB.Model(&model.Task{}).
		Preload("User").
		Where("uid=?", uid).
		Where("title LIKE ? OR content LIKE ?", "%"+service.Info+"%", "%"+service.Info+"%"). //模糊查询
		Count(&cnt).
		Limit(service.PageSize).Offset((service.PageNum - 1) * (service.PageSize)).
		Find(&tasks).Error
	if err != nil {
		return serializer.Response{
			Status:  400,
			Message: "查询所有已完成事项失败",
		}
	} else {
		return serializer.Response{
			Status:  200,
			Data:    serializer.BuildListResponse(serializer.Build_Finished_Tasks(tasks), uint(cnt)),
			Message: "查询所有已完成事项成功",
		}
	}

}
func (service *SearchtaskService) Show_UnFinished(uid uint) serializer.Response {
	var tasks []model.Task
	cnt := 0
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	err := model.DB.Model(&model.Task{}).
		Preload("User").
		Where("uid=?", uid).
		Where("title LIKE ? OR content LIKE ?", "%"+service.Info+"%", "%"+service.Info+"%"). //模糊查询
		Count(&cnt).
		Limit(service.PageSize).Offset((service.PageNum - 1) * (service.PageSize)).
		Find(&tasks).Error
	if err != nil {
		return serializer.Response{
			Status:  400,
			Message: "查询所有未完成事项失败",
		}
	} else {
		return serializer.Response{
			Status:  200,
			Data:    serializer.BuildListResponse(serializer.Build_UnFinished_Tasks(tasks), uint(cnt)),
			Message: "查询所有未完成事项成功",
		}
	}

}
func (service *SearchtaskService) Show_All(uid uint) serializer.Response {
	var tasks []model.Task
	cnt := 0
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	err := model.DB.Model(&model.Task{}).
		Preload("User").
		Where("uid=?", uid).
		Where("title LIKE ? OR content LIKE ?", "%"+service.Info+"%", "%"+service.Info+"%"). //模糊查询
		Count(&cnt).
		Limit(service.PageSize).Offset((service.PageNum - 1) * (service.PageSize)).
		Find(&tasks).Error
	if err != nil {
		return serializer.Response{
			Status:  400,
			Message: "查询所有事项失败",
		}
	} else {
		return serializer.Response{
			Status:  200,
			Data:    serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(cnt)),
			Message: "查询所有事项成功",
		}
	}

}
