package router

import (
	"github.com/gin-gonic/gin"
	"go_crud_demo/controllers"
)

func Router() *gin.Engine{
	r := gin.Default()

	// 以中间件的形式在路由中调用logger
	//r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	//r.Use(logger.Recover)


    // user router
	userRouter := r.Group("/user")
	// router use controllers   // 大项目推荐用下面这种方式,以防方法名冲突
	{
		//userRouter.POST("/add", controllers.UserController{}.AddUserInfo)  // todo 需要先改好sqldb，增加model

		userRouter.POST("/upload", controllers.UserController{}.UploadFile)   // 20240514 add upload api, and migrate to router
		userRouter.GET("/download", controllers.UserController{}.DownloadFile) 	// for test vue download file code

	}

	return r
}