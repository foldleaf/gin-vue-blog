package routes

import (
	v1 "gin-vue-blog/api/v1"
	"gin-vue-blog/utils"
	//"net/http"

	"github.com/gin-gonic/gin"
	//v1 "go.opentelemetry.io/proto/otlp/common/v1"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	//v1路由组
	router := r.Group("api/v1")
	{
		//Use用户模块路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)

		//Category分类模块路由接口

		//Article文章模块路由接口
	}

	r.Run(utils.HttpPort)
}
