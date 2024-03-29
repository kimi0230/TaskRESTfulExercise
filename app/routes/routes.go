package routes

import (
	"TaskRESTfulExercise/app/api/heartbeat"
	"TaskRESTfulExercise/app/api/task"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/ping", heartbeat.Ping)

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/tasks", task.PostTasks)
			v1.GET("/tasks", task.GetTasks)
			v1.PUT("/tasks/:id", task.PutTasks)
			v1.DELETE("/tasks/:id", task.DeleteTasks)
		}

	}
}
