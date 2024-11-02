package router

import (
	v1 "MiMengCore/api/v1"
	"github.com/gin-gonic/gin"
)

//暂时用不上
/*
func InitRouter(Router *gin.RouterGroup) {
	NoticeRouter := Router.Group("/notice")
	{
		NoticeRouter.GET("", v1.GetNotice)
		NoticeRouter.PUT("", v1.UpdateNotice)
	}
}
*/

func InitRouter(Router *gin.RouterGroup) {
	ContentRouter := Router.Group("/content")
	{
		ContentRouter.GET("/notice", v1.GetNoticeHandler)
		ContentRouter.PUT("/notice", v1.UpdateNoticeHandler)
	}

	UserRouter := Router.Group("/user")
	{
		UserRouter.POST("", v1.Register)
	}

	AuthRouter := Router.Group("/auth")
	{
		AuthRouter.POST("/login", v1.Login)
	}
}
