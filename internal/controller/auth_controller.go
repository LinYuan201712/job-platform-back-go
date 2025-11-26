package controller

import (
	"github.com/gin-gonic/gin"
	"job-platform-go/internal/model/dto"
	"job-platform-go/internal/service"
	"job-platform-go/pkg/e"
	"job-platform-go/pkg/response"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		authService: service.NewAuthService(),
	}
}
func (ctrl *AuthController) Login(c *gin.Context) {
	var req dto.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorWithStatus(c, 400, e.ERROR_BAD_REQUEST, "请求参数错误")
		return
	}
	resp, err := ctrl.authService.Login(req)
	if err != nil {
		// 业务错误通常返回 200 + 错误码，或者 401
		response.Error(c, 401, err.Error())
		return
	}
	response.Success(c, resp)

}

func (ctrl *AuthController) Register(c *gin.Context) {
	var req dto.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorWithStatus(c, 400, e.ERROR_BAD_REQUEST, "请求参数错误")
		return
	}
	err := ctrl.authService.Register(req)
	if err != nil {
		response.Error(c, 401, err.Error())
		return
	}
	msg := "学生账户注册成功"
	code := 201
	if req.Role == "hr" {
		msg = "企业账户注册成功，请等待管理员审核"
		code = 202
	}

	// 这里手动构建 Response 以支持不同的 Code
	c.JSON(200, map[string]interface{}{
		"code":    code,
		"message": msg,
		"data":    nil,
	})
}
