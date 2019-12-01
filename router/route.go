package router

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yuzhikuan/user_manage/controller"
	"github.com/yuzhikuan/user_manage/middleware"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)

	store := sessions.NewCookieStore([]byte("secret"))
	apiNormalGroup := router.Group("/api")
	apiController := &controller.Api{}
	apiNormalGroup.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.TranslationMiddleware(),
	)
	apiNormalGroup.POST("/login", apiController.Login)
	apiNormalGroup.GET("/loginout", apiController.LoginOut)

	apiAuthGroup := router.Group("/api")
	apiAuthGroup.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.SessionAuthMiddleware(),
		middleware.TranslationMiddleware(),
	)
	apiAuthGroup.GET("/user/listpage", apiController.ListPage)
	apiAuthGroup.GET("/user/add", apiController.AddUser)
	apiAuthGroup.GET("/user/edit", apiController.EditUser)
	apiAuthGroup.GET("/user/remove", apiController.RemoveUser)
	apiAuthGroup.GET("/user/batchremove", apiController.RemoveUser)
	return router
}
