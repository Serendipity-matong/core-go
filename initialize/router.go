package initialize

import (
	v1 "MiMengCore/api/v1"
	"MiMengCore/router"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//r.Use(middleware.Cors())
	r.GET("/", v1.Index)
	Group := r.Group("api/v1")
	{

		router.InitRouter(Group)
		//Group.GET("/notice", v1.GetNotice)
		//Group.PUT("/notice", v1.UpdateNotice)
	}
	Group2 := r.Group("api/v2")
	{

		router.InitRouter2(Group2)
		//Group.GET("/notice", v1.GetNotice)
		//Group.PUT("/notice", v1.UpdateNotice)
	}
	return r
}
