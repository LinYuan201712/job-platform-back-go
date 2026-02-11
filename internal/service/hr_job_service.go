package service

import (
	"github.com/pkg/errors"
	"job-platform-go/internal/model/dto"
	"job-platform-go/internal/model/entity"
	"job-platform-go/internal/repository"
	"time"
)

type HrJobService struct {
	jobRepo     *repository.JobRepository
	companyRepo *repository.CompanyRepository
}

func NewHrJobService() *HrJobService {
	return &HrJobService{
		jobRepo:     repository.NewJobRepository(),
		companyRepo: repository.NewCompanyRepository(),
	}
}

// 发布岗位
func (s *HrJobService) CreateJob(userID int, req dto.HrJobCreateReq) error {
	//1.由userID查company
	company, err := s.companyRepo.FindByUserID(userID)
	if err != nil {
		return errors.New("企业不存在")
	}
	//2. 参数映射
	job := &entity.Job{
		CompanyID:        company.CompanyID,
		PostedByUserID:   userID,
		Title:            req.Title,
		Department:       req.Department,
		Description:      req.Description,
		TechRequirements: req.TechRequirements,
		BonusPoints:      req.BonusPoints,
		MinSalary:        req.MinSalary,
		MaxSalary:        req.MaxSalary,
		ProvinceID:       &req.ProvinceID,
		CityID:           &req.CityID,
		AddressDetail:    req.AddressDetail,
		WorkAddress:      req.WorkAddress,
		Headcount:        req.Headcount,
		Type:             &req.Type,
		RequiredDegree:   req.RequiredDegree,
	}
	s.fillJobEnumsAndDates(job, req.WorkNature, req.Status, req.Deadline, req.RequiredStartDate)
	//3.调用repo插入
	err = s.jobRepo.CreateJob(job, req.Tags)
	if err != nil {
		return errors.New("发布岗位失败")
	}
	return nil
}

// 更新岗位
func (s *HrJobService) UpdateJob(userID int, jobID int, req dto.HrJobUpdateReq) error {
	//1.由userID查company
	company, err := s.companyRepo.FindByUserID(userID)
	if err != nil {
		return errors.New("企业不存在")
	}
	//岗位是否存在
	//校验权限：只能改自己公司的职位
	job, err := s.jobRepo.FindByIDAndCompany(jobID, company.CompanyID)
	if err != nil {
		return errors.New("岗位不存在或无权操作")
	}
	//2. 参数映射
	job.Title = req.Title
	job.Department = req.Department
	job.Description = req.Description
	job.TechRequirements = req.TechRequirements
	job.BonusPoints = req.BonusPoints
	job.MinSalary = req.MinSalary
	job.MaxSalary = req.MaxSalary
	job.ProvinceID = &req.ProvinceID
	job.CityID = &req.CityID
	job.AddressDetail = req.AddressDetail
	job.WorkAddress = req.WorkAddress
	job.Headcount = req.Headcount
	job.Type = &req.Type
	job.RequiredDegree = req.RequiredDegree

	s.fillJobEnumsAndDates(job, req.WorkNature, req.Status, req.Deadline, req.RequiredStartDate)
	//3.调用repo插入
	err = s.jobRepo.Update(job, req.Tags)
	if err != nil {
		return errors.New("关系岗位失败")
	}
	return nil
}

/*
//下线岗位
CloseJob
//岗位列表
ListJobs
//获取岗位详情
GetJobDetail

*/
// 辅助方法：填充枚举和日期
func (s *HrJobService) fillJobEnumsAndDates(job *entity.Job, workNatureStr, statusStr, deadlineStr, startDateStr string) {
	wn := entity.WorkNatureInternship
	if workNatureStr == "full-time" || workNatureStr == "校招" {
		wn = entity.WorkNatureFullTime
	}
	job.WorkNature = &wn

	st := entity.JobStatusDraft
	if statusStr == "pending" {
		st = entity.JobStatusPending
	}
	if statusStr == "approved" {
		st = entity.JobStatusApproved
	}
	job.Status = st

	if t, err := time.Parse("2006-01-02", deadlineStr); err == nil {
		job.Deadline = &t
	}
	if t, err := time.Parse("2006-01-02", startDateStr); err == nil {
		job.RequiredStartDate = &t
	}
}
