package controller

import (
	"github.com/gin-gonic/gin"
	"job-platform-go/internal/service"
	"job-platform-go/pkg/response"
)

type CommonController struct {
	locService *service.LocationService
	tagService *service.TagService
}

func NewCommonController() *CommonController {
	return &CommonController{
		locService: service.NewLocationService(),
		tagService: service.NewTagService(),
	}
}

func (ctrl *CommonController) GetLocations(c *gin.Context) {
	data, err := ctrl.locService.GetProvincesCities()
	if err != nil {
		response.Error(c, 500, "获取省市数据失败")
		return
	}
	response.Success(c, data)
}

func (ctrl *CommonController) GetAllTags(c *gin.Context) {
	data, err := ctrl.tagService.GetAllGroupedTags()
	if err != nil {
		response.Error(c, 500, "获取标签数据失败")
		return
	}
	response.Success(c, data)
}
