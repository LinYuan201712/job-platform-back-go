package service

import (
	"errors"
	"job-platform-go/internal/model/dto"
	"job-platform-go/internal/model/entity"
	"job-platform-go/internal/repository"
)

type CompanyService struct {
	companyRepo    *repository.CompanyRepository
	dictionaryRepo *repository.DictionaryRepository
}

func NewCompanyService() *CompanyService {
	return &CompanyService{
		companyRepo:    repository.NewCompanyRepository(),
		dictionaryRepo: repository.NewDictionaryRepository(),
	}
}

// GetProfile 获取企业详情
func (s *CompanyService) GetProfile(userID int) (*dto.CompanyProfileResp, error) {
	//1. 查企业基本信息
	company, err := s.companyRepo.FindByUserID(userID)
	if err != nil {
		return nil, errors.New("查找企业基本信息失败")
	}
	//2. 查链接
	links, err := s.companyRepo.GetLinks(company.CompanyID)
	if err != nil {
		return nil, errors.New("查找企业链接失败")
	}
	//查字典
	var industryName, natureName, scaleName string

	if company.IndustryID != nil {
		name, _ := s.dictionaryRepo.GetIndustryNameByID(*company.IndustryID)
		industryName = name
	}
	if company.NatureID != nil {
		name, _ := s.dictionaryRepo.GetNatureNameByID(*company.NatureID)
		natureName = name
	}
	if company.CompanyScaleID != nil {
		name, _ := s.dictionaryRepo.GetScaleNameByID(*company.CompanyScaleID)
		scaleName = name
	}
	//3. 组装响应
	resp := &dto.CompanyProfileResp{
		CompanyName:        company.CompanyName,
		Description:        company.Description,
		LogoUrl:            company.LogoUrl,
		Nature:             natureName,
		Industry:           industryName,
		CompanyScale:       scaleName,
		ContactPersonName:  company.ContactPersonName,
		ContactPersonPhone: company.ContactPersonPhone,
		CompanyAddress:     company.CompanyAddress,
		// 统计数据 (暂时 Mock，后续对接 Job 模块时再完善)
		OpenJobsCount:     0,
		ResumeProcessRate: "0",
		ExternalLinks:     make([]dto.ExternalLinkDto, 0),
	}

	//4. 转换链接
	for _, l := range links {
		resp.ExternalLinks = append(resp.ExternalLinks, dto.ExternalLinkDto{
			LinkName: l.LinkName,
			LinkUrl:  l.LinkUrl,
		})
	}
	return resp, nil

}

// UpdateProfile ()更新企业详情
func (s *CompanyService) UpdateProfile(userID int, req dto.CompanyProfileUpdateReq) error {
	company, err := s.companyRepo.FindByUserID(userID)
	if err != nil {
		return errors.New("企业不存在")
	}
	company.Description = req.Description
	company.CompanyAddress = req.CompanyAddress
	company.ContactPersonName = req.ContactPersonName
	company.ContactPersonPhone = req.ContactPersonPhone
	//查字典
	industryID, err := s.dictionaryRepo.FindIndustryIDByName(req.Industry)
	if err != nil {
		return errors.New("行业不存在")
	}
	natureID, err := s.dictionaryRepo.FindNatureIDByName(req.Nature)
	if err != nil {
		return errors.New("性质不存在")
	}
	scaleID, err := s.dictionaryRepo.FindScaleIDByName(req.CompanyScale)
	if err != nil {
		return errors.New("规模不存在")
	}

	company.NatureID = &natureID
	company.IndustryID = &industryID
	company.CompanyScaleID = &scaleID

	// 构建 Links 实体列表
	var links []entity.CompanyExternalLink
	for _, l := range req.ExternalLinks {
		links = append(links, entity.CompanyExternalLink{
			CompanyID: company.CompanyID,
			LinkName:  l.LinkName,
			LinkUrl:   l.LinkUrl,
		})
	}
	return s.companyRepo.UpdateProfileWithLinks(company, links)

}
