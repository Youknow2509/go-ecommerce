package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`    // Ma status code
	Message string      `json:"message"` // Thong bao loi
	Data    interface{} `json:"data"`    // Du lieu duoc return
}

type ErrResponseData struct {
	Code   int         `json:"code"`   // Ma status code
	Error  string      `json:"error"`  //
	Detail interface{} `json:"detail"` // Thong bao loi
}

// func success response
func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

// func error response
func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}
