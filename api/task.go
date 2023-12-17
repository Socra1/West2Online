package api

import (
	"fmt"
	"net/http"
	"todo_list/service"
	utils "todo_list/token"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var createTask service.Create_UpdateTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	c.Bind(&createTask)
	res := createTask.Create(claim.Id)
	fmt.Printf("res: %v\n", res)
	c.JSON(http.StatusOK, gin.H{
		"Status":  res.Status,
		"Message": res.Message,
		"Data":    res.Data,
		"Error":   res.Error,
	})
}

func SearchTask(c *gin.Context) {
	var searchtask service.SearchtaskService
	c.Bind(&searchtask)
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	res := searchtask.Search(claim.Id)
	c.JSON(http.StatusOK, gin.H{
		"Status":  res.Status,
		"Message": res.Message,
		"Data":    res.Data,
		"Error":   res.Error,
	})
}

func FinishATask(c *gin.Context) {
	var Atask service.Un_or_FinishiATaskService
	c.Bind(&Atask)
	res := Atask.Un_or_FinishOne(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"Status":  res.Status,
		"Message": res.Message,
		"Data":    res.Data,
		"Error":   res.Error,
	})
}

func FinishAllTask(c *gin.Context) {
	var Alltask service.Un_or_FinishiAllTaskService
	c.Bind(&Alltask)
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	res := Alltask.Un_or_FinishAll(claim.Id)
	c.JSON(http.StatusOK, gin.H{
		"Status":  res.Status,
		"Message": res.Message,
		"Data":    res.Data,
		"Error":   res.Error,
	})
}

func ShowFinished(c *gin.Context) {
	var finishedtask service.SearchtaskService
	c.Bind(&finishedtask)
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	res := finishedtask.Show_Finished(claim.Id)
	c.JSON(http.StatusOK, gin.H{
		"Status":  res.Status,
		"Message": res.Message,
		"Data":    res.Data,
		"Error":   res.Error,
	})
}

func ShowUnFinished(c *gin.Context) {
	var unfinishedtask service.SearchtaskService
	c.Bind(&unfinishedtask)
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	res := unfinishedtask.Show_UnFinished(claim.Id)
	c.JSON(http.StatusOK, gin.H{
		"Status":  res.Status,
		"Message": res.Message,
		"Data":    res.Data,
		"Error":   res.Error,
	})
}

func ShowAll(c *gin.Context) {
	var alltasks service.SearchtaskService
	c.Bind(&alltasks)
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	res := alltasks.Show_All(claim.Id)
	c.JSON(http.StatusOK, gin.H{
		"Status":  res.Status,
		"Message": res.Message,
		"Data":    res.Data,
		"Error":   res.Error,
	})
}

func DeleteTask(c *gin.Context) {
	var deletetask service.EmptyTask
	c.Bind(&deletetask)
	res := deletetask.Delete(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"Status":  res.Status,
		"Message": res.Message,
		"Data":    res.Data,
		"Error":   res.Error,
	})
}
