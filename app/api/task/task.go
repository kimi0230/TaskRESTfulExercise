package task

import (
	"TaskRESTfulExercise/app/models/taskmodel"
	"TaskRESTfulExercise/configs/errorCode"
	"TaskRESTfulExercise/services/ginservices"
	"TaskRESTfulExercise/services/mongodb"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func PostTasks(c *gin.Context) {
	type structRequest struct {
		Name        string `json:"name" form:"name" binding:"required,max=30"`
		Description string `json:"description" form:"description" binding:"required,max=256"`
		Priority    int    `json:"priority" form:"priority" binding:"number"`
		DueDate     string `json:"dueDate,omitempty" form:"dueDate,omitempty" binding:"-" time_format:"2006-01-02" time_utc:"8"`
	}
	var reqJSON structRequest
	_, err := ginservices.GinRequest(c, &reqJSON)
	if err != nil {
		ginservices.GinRespone(c, "", "", errorCode.PARAMS_INVALID, err)
		return
	}

	// Create mongo client
	mgClient, err := mongodb.NewMongoClient()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		ginservices.GinRespone(c, "", "", errorCode.INTERAL_SERVER_ERROR, err)
		return
	}
	taskDAO := taskmodel.NewDAO(mgClient)

	newTask := &taskmodel.TaskDTO{
		Name:        reqJSON.Name,
		Description: reqJSON.Description,
		Priority:    reqJSON.Priority,
		DueDate:     reqJSON.DueDate,
	}

	result, err := taskDAO.Create(newTask)
	if err != nil {
		ginservices.GinRespone(c, "", "", errorCode.INTERAL_SERVER_ERROR, err)
		return
	}

	ginservices.GinRespone(c, "", result, errorCode.SUCCESS, nil)
}

func GetTasks(c *gin.Context) {
	type structRequest struct {
		taskmodel.PageQueryArgs
	}
	var reqJSON structRequest
	_, err := ginservices.GinRequest(c, &reqJSON)
	if err != nil {
		ginservices.GinRespone(c, "", "", errorCode.PARAMS_INVALID, err)
		return
	}

	// Create mongo client
	mgClient, err := mongodb.NewMongoClient()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		ginservices.GinRespone(c, "", "", errorCode.INTERAL_SERVER_ERROR, err)
		return
	}
	taskDAO := taskmodel.NewDAO(mgClient)
	skip := (reqJSON.Page - 1) * reqJSON.Limit

	filter := bson.M{
		"$or": []bson.M{
			// "i" 表示忽略大小寫
			{"name": bson.M{"$regex": primitive.Regex{Pattern: reqJSON.Keyword, Options: "i"}}},
		},
	}

	result, _ := taskDAO.GetByQuery(filter, int64(reqJSON.Limit), int64(skip), reqJSON.Order, reqJSON.By)

	if len(*result) == 0 {
		ginservices.GinRespone(c, "", "", errorCode.NO_CONTENT, nil)
		return
	}
	total, err := taskDAO.CountDocuments(filter)
	if err != nil {
		ginservices.GinRespone(c, "", "", errorCode.INTERAL_SERVER_ERROR, err)
		return
	}

	c.Writer.Header().Set("x-page", strconv.Itoa(reqJSON.Page))
	c.Writer.Header().Set("x-per-page", strconv.Itoa(reqJSON.Limit))
	c.Writer.Header().Set("x-total", strconv.Itoa(int(total)))
	totalPages := math.Ceil(float64(total) / float64(reqJSON.Limit))
	c.Writer.Header().Set("x-total-pages", strconv.Itoa(int(totalPages)))

	ginservices.GinRespone(c, "", result, errorCode.SUCCESS, nil)
}