package router

import (
	"github.com/gin-gonic/gin"
	apiv1auth "github.com/water25234/golang-postgres/app/controller/api/v1/auth"
	apiv1user "github.com/water25234/golang-postgres/app/controller/api/v1/user"
	"github.com/water25234/golang-postgres/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	v1 := router.Group("/api/v1")
	{
		v1.Use(middleware.ExecuteThrottle())
		authRouting := v1.Group("/auth")
		{
			authRouting.GET("", apiv1auth.GetAuth)

			authRouting.GET("/throttle", apiv1auth.GetThrottle)

			authRouting.DELETE("/:id", apiv1auth.DeleteAuth)

			authRouting.POST("/:uid", apiv1auth.PostAuth)
		}

		userRouting := v1.Group("/user")
		{
			userRouting.GET("/list", apiv1user.GetUserListVersion2)

			userRouting.GET("/item/:uid", apiv1user.GetUser)

			userRouting.POST("/addUser", apiv1user.AddUser)
		}
	}

	return router
}
