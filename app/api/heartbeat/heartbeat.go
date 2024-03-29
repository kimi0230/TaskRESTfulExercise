package heartbeat

import (
	"TaskRESTfulExercise/configs/errorCode"
	"TaskRESTfulExercise/services/ginservices"
	"TaskRESTfulExercise/services/mongodb"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Ping(c *gin.Context) {
	if _, err := mongodb.ConnectMongoDB(viper.GetString("mongo.ip"), viper.GetString("mongo.port"), viper.GetString("mongo.username"), viper.GetString("mongo.password"), viper.GetString("mongo.poolsize"), viper.GetString("mongo.database")); err != nil {
		ginservices.GinRespone(c, "", "", errorCode.INTERAL_SERVER_ERROR, err)
		return
	}
	ginservices.GinRespone(c, "", "pong", errorCode.SUCCESS, nil)
}
