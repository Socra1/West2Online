package serializer

import (
	"todo_list/model"
)

type Task struct {
	ID        uint   `json:"id" example:"1"`       //任务id
	Title     string `json:"title" example:"吃饭"`   //题目
	Content   string `json:"content" example:"睡觉"` //内容
	View      uint64 `json:"view" example:"0"`     //浏览量
	Status    int    `json:"status" example:""`    //0为未完成
	CreateAt  int64  `json:"create_at"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time" `
}

// 任务序列化
func BuildTask(item model.Task) Task {
	return Task{
		ID:        item.ID,
		Title:     item.Title,
		Content:   item.Content,
		Status:    item.Status,
		StartTime: item.StartTime,
		EndTime:   item.EndTime,
	}
}
func Build_Finished_Tasks(items []model.Task) (tasks []Task) {
	for _, i := range items {
		task := BuildTask(i)
		if task.Status == 1 {
			tasks = append(tasks, task)
		}
	}
	return tasks
}
func Build_UnFinished_Tasks(items []model.Task) (tasks []Task) {
	for _, i := range items {
		task := BuildTask(i)
		if task.Status == 0 {
			tasks = append(tasks, task)
		}
	}
	return tasks
}
func BuildTasks(items []model.Task) (tasks []Task) {
	for _, i := range items {
		task := BuildTask(i)
		tasks = append(tasks, task)
	}
	return tasks
}
