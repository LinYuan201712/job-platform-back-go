package controller

import (
	"github.com/gin-gonic/gin"
	"job-platform-go/internal/model/dto"
	"job-platform-go/internal/service"
	"job-platform-go/pkg/e"
	"job-platform-go/pkg/response"
)

type CompanyController struct {
	companyService *service.CompanyService
}

func NewCompanyController() *CompanyController {
	return &CompanyController{
		companyService: service.NewCompanyService(),
	}
}

func (ctrl *CompanyController) GetProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.Error(c, 401, "用户未登录")
		return
	}

	resq, err := ctrl.companyService.GetProfile(userID.(int))
	if err != nil {
		response.Error(c, 401, err.Error())
		return
	}
	response.Success(c, resq)

}

func (ctrl *CompanyController) UpdateProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.Error(c, 401, "用户未登录")
		return
	}
	var req dto.CompanyProfileUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorWithStatus(c, 400, e.ERROR_BAD_REQUEST, "请求参数错误")
		return
	}
	err := ctrl.companyService.UpdateProfile(userID.(int), req)
	if err != nil {
		response.Error(c, 401, err.Error())
		return
	}
	response.Success(c, nil)

}
