package ginservices

import (
	"bytes"
	"io"

	"github.com/gin-gonic/gin"
)

type ResponseObj struct {
	Data  interface{} `json:"data"`
	Code  string      `json:"code"`
	Error string      `json:"error"`
}

func GinRequest(c *gin.Context, reqJSON interface{}) (interface{}, error) {
	var reqData interface{}
	if c.Request.Method == "GET" {
		reqData = c.Request.URL.Query()
		if bindErr := c.ShouldBind(reqJSON); bindErr != nil {
			return nil, bindErr
		}
	} else {
		data, err := c.GetRawData()
		if err != nil {
			return nil, err
		}
		reqData = io.NopCloser(bytes.NewBuffer(data))

		c.Request.Body = io.NopCloser(bytes.NewBuffer(data))
		if bindErr := c.ShouldBind(reqJSON); bindErr != nil {
			return nil, bindErr
		}

	}
	return reqData, nil
}

func GinRespone(c *gin.Context, resquestData interface{}, responseData interface{}, resultMsg map[string]interface{}, err interface{}) {
	var response ResponseObj
	if responseData == "" || responseData == nil {
		response.Data = nil
	} else {
		response.Data = responseData
	}

	response.Code = resultMsg["code"].(string)

	if err != "" && err != nil {
		response.Error = err.(error).Error()
	}

	c.JSON(resultMsg["httpCode"].(int), response)
}
