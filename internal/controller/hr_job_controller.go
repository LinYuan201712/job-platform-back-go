package controller

import (
	"github.com/gin-gonic/gin"
	"job-platform-go/internal/model/dto"
	"job-platform-go/internal/service"
	"job-platform-go/pkg/response"
	"strconv"
)

type HrJobController struct {
	hrJobService *service.HrJobService
}

func NewHrJobController() *HrJobController {
	return &HrJobController{
		hrJobService: service.NewHrJobService(),
	}
}

// 发布岗位
func (ctr *HrJobController) CreateJob(c *gin.Context) {
	//json绑定
	var req dto.HrJobCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "参数错误")
		return
	}
	userID := c.GetInt("userID")
	//调用service
	if err := ctr.hrJobService.CreateJob(userID, req); err != nil {
		response.Error(c, 500, err.Error())
		return
	}
	response.Success(c, nil)
}

// 更新岗位
func (ctr *HrJobController) UpdateJob(c *gin.Context) {
	jobID, _ := strconv.Atoi(c.Param("id"))
	//json绑定
	var req dto.HrJobUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "参数错误")
		return
	}
	userID := c.GetInt("userID")
	//调用service
	if err := ctr.hrJobService.UpdateJob(userID, jobID, req); err != nil {
		response.Error(c, 500, err.Error())
		return
	}
	response.Success(c, nil)
}
