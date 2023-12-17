package routers

import (
	"todo_list/api"
	"todo_list/midddleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))

	r.Use(sessions.Sessions("mysession", store))
	r.GET("/swagger/*any", ginswagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("api/v1")
	{
		v1.POST("/register", api.UserRegisrter)
		v1.POST("/login", api.UserLogin)
		//进行身份认证（中间件认证）
		authed := v1.Group("/")
		authed.Use(midddleware.JWT())
		{
			//增
			authed.POST("/tasks", api.CreateTask)
			//查
			authed.GET("/finishedtasks", api.ShowFinished)
			authed.GET("/unfinishedtasks", api.ShowUnFinished)
			authed.GET("/tasks", api.ShowAll)
			authed.POST("/search", api.SearchTask)

			//改
			authed.PUT("/changestates/:id", api.FinishATask)
			authed.PUT("/changestates", api.FinishAllTask)

			//删
			authed.DELETE("/tasks/:id", api.DeleteTask)
		}
	}
	return r
}
