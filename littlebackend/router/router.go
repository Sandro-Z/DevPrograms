package router

import (
	"exam_go/controller"
	"exam_go/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// r.Use(middleware.Error)
	apiRouter := r.Group("/api")
	{
		helloController := controller.HelloController{}
		apiRouter.GET("", helloController.Hello)
		apiRouter.GET("/time", helloController.HelloTime)

		userController := controller.UserController{}
		{
			apiRouter.GET("/userinfo", middleware.IsLogin, userController.HandleUserinfo)
			apiRouter.POST("/register", userController.HandleSignup)
			apiRouter.POST("/login", userController.HandleLogin)
			apiRouter.DELETE("/logout", userController.HandleLogout)
			apiRouter.GET("/getRestaurants", middleware.IsLogin, userController.HandleGetRestList)
			apiRouter.GET("/foods", middleware.IsLogin, userController.HandleGetRestFoodList)
			apiRouter.GET("/food/:id", middleware.IsLogin, userController.HandleGetFoodInfo)
			apiRouter.POST("/food/select", middleware.IsLogin, userController.HandleOrderFood)
			apiRouter.GET("/foods/status/:id", middleware.IsLogin, userController.HandleOrderInfo)
			apiRouter.POST("/foods/like", middleware.IsLogin, userController.HandleLikeInfo)
			// apiRouter.POST("/login", middleware.IsNotLogin, userController.HandleLogin)
		}

		restRouter := apiRouter.Group("restaurant")
		{
			restController := controller.RestController{}
			restRouter.GET("/info", middleware.IsRest, restController.HandleRestinfo)
			restRouter.POST("/food/add", middleware.IsRest, restController.HandleRestFoodAdd)
			restRouter.POST("/apply", middleware.IsLogin, restController.HandleNewRest)
			restRouter.GET("/checkStatus", middleware.IsRest, restController.HandleNewRestinfo)
			restRouter.GET("/checkAll", middleware.IsRest, restController.HandleCheck)
			restRouter.POST("/changeStatus", middleware.IsRest, restController.HandleChange)
			restRouter.POST("/food/change", middleware.IsRest, restController.HandleChangeFood)
		}

		adminRouter := apiRouter.Group("admin")
		{
			adminController := controller.AdminController{}
			adminRouter.GET("", middleware.IsAdmin, adminController.AdminHello)
			adminRouter.GET("/restaurant", middleware.IsAdmin, adminController.AdminRestList)
			adminRouter.GET("/foods", middleware.IsAdmin, adminController.AdminRestFoodList)
			adminRouter.POST("/check", middleware.IsAdmin, adminController.AdminChange)
		}
	}
}
