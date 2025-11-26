package response

import (
	"github.com/gin-gonic/gin"
	"job-platform-go/pkg/e"
	"net/http"
)

// 分页结构
type Pagination struct {
	TotalItems  int64 `json:"total_items"`
	TotalPages  int   `json:"total_pages"`
	CurrentPage int   `json:"current_page"`
	PageSize    int   `json:"page_size"`
}

type Response struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination,omitempty"` //值为空值时，不输出到 JSON 中
}

// Result() —— 统一出口方法
func Result(c *gin.Context, httpStatus int, code int, msg string, data interface{}, pagination *Pagination) {
	c.JSON(httpStatus, Response{
		Code:       code,
		Message:    msg,
		Data:       data,
		Pagination: pagination,
	})
}

// Success 成功响应 (HTTP 200)
func Success(c *gin.Context, data interface{}) {
	Result(c, http.StatusOK, e.SUCCESS, e.GetMsg(e.SUCCESS), data, nil)
}

// Success 成功响应 (HTTP 200)带分页
func SuccessWithPage(c *gin.Context, data interface{}, pagination Pagination) {
	Result(c, http.StatusOK, e.SUCCESS, e.GetMsg(e.SUCCESS), data, &pagination)
}

// Error 错误响应
func Error(c *gin.Context, code int, msg string) {
	if msg == "" {
		msg = e.GetMsg(code)
	}
	// 默认使用 HTTP 200，通过 code 判断业务错误
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}
func ErrorWithStatus(c *gin.Context, httpStatus int, code int, msg string) {
	if msg == "" {
		msg = e.GetMsg(code)
	}
	// 指定 HTTP 状态码的错误响应
	c.JSON(httpStatus, Response{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}
