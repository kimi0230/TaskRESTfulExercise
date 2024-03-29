package routes

import (
	"TaskRESTfulExercise/app/api/heartbeat"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/ping", heartbeat.Ping)

	// api := router.Group("/api")
	// {
	// 	v1 := api.Group("/v1")
	// 	{

	// 	}

	// }
}
