package errorCode

import "net/http"

/*
code :Http `Error Code`-`分類`-`版本`-`編號`
*/

/* 2xx Successful responses */
// 200
var SUCCESS = map[string]interface{}{
	"code":     "200-API-V1-0001",
	"status":   true,
	"message":  "Success",
	"httpCode": http.StatusOK,
}

// 202
var ACCEPTED = map[string]interface{}{
	"code":     "202-API-V1-0001",
	"status":   false,
	"message":  "Accepted",
	"httpCode": http.StatusAccepted,
}

// 204
var NO_CONTENT = map[string]interface{}{
	"code":     "204-API-V1-0001",
	"status":   false,
	"message":  "no content",
	"httpCode": http.StatusNoContent,
}

/******/

/* 3xx Redirection messages*/

/******/

/* 4xx Client error responses ( */
// 400-API-V1-0001 StatusBadRequest
var BAD_REQUEST = map[string]interface{}{
	"code":     "400-API-V1-0001",
	"status":   false,
	"message":  "Bad Request",
	"httpCode": http.StatusBadRequest,
}

// 400-API-V1-0002 StatusBadRequest
var PARAMS_INVALID = map[string]interface{}{
	"code":     "400-API-V1-0002",
	"status":   false,
	"message":  "必填、不可填的參數有誤",
	"httpCode": http.StatusBadRequest,
}

// 400-API-V1-0003 StatusBadRequest
var CAPTCHA_INVALID = map[string]interface{}{
	"code":     "400-API-V1-0003",
	"status":   false,
	"message":  "驗證碼錯誤or過期",
	"httpCode": http.StatusBadRequest,
}

// 401 StatusUnauthorized
var UNAUTHORIZED = map[string]interface{}{
	"code":     "401-API-V1-0001",
	"status":   false,
	"message":  "Forbidden Request",
	"httpCode": http.StatusUnauthorized,
}

// 401 StatusUnauthorized
var HEADER_NO_AUTHORIZATION = map[string]interface{}{
	"code":     "401-API-V1-0002",
	"status":   false,
	"message":  "Header No authorization",
	"httpCode": http.StatusUnauthorized,
}

// 403 StatusForbidden
var FORBIDDEN = map[string]interface{}{
	"code":     "403-API-V1-0001",
	"status":   false,
	"message":  "Forbidden",
	"httpCode": http.StatusForbidden,
}

// 406 StatusNotAcceptable
var NOT_ACCEPTABLE = map[string]interface{}{
	"code":     "406-API-V1-0001",
	"status":   false,
	"message":  "Not Acceptable",
	"httpCode": http.StatusNotAcceptable,
}

// 408 StatusRequestTimeout
var REQUEST_TIMEOUT = map[string]interface{}{
	"code":     "408-API-V1-0001",
	"status":   false,
	"message":  "Timeout or Duplicate",
	"httpCode": http.StatusRequestTimeout,
}

// 429 StatusTooManyRequests
var RATE_LIMIT = map[string]interface{}{
	"code":     "429-API-V1-0001",
	"status":   false,
	"message":  "Too many requests",
	"httpCode": http.StatusTooManyRequests,
}

/******/

/* 5xx Server error responses */
// 500-API-V1-0001 StatusInternalServerError
var INTERAL_SERVER_ERROR = map[string]interface{}{
	"code":     "500-API-V1-0001",
	"status":   false,
	"message":  "Server",
	"httpCode": http.StatusInternalServerError,
}

// 500-API-V1-0002 StatusInternalServerError
var DBUPDATEFAIL = map[string]interface{}{
	"code":     "500-API-V1-0002",
	"status":   false,
	"message":  "DB write error",
	"httpCode": http.StatusInternalServerError,
}

// 500-API-V1-0003 StatusInternalServerError
var DBQUERYFAIL = map[string]interface{}{
	"code":     "500-API-V1-0003",
	"status":   false,
	"message":  "DB query error",
	"httpCode": http.StatusInternalServerError,
}

// 500-API-V1-0004 StatusInternalServerError
var AWS_S3_ERROR = map[string]interface{}{
	"code":     "500-API-V1-0004",
	"status":   false,
	"message":  "AWS S3",
	"httpCode": http.StatusInternalServerError,
}

/******/
