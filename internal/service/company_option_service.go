package service

import (
	"job-platform-go/internal/model/vo"
	"job-platform-go/internal/repository"
)

type CompanyOptionService struct {
	repo *repository.DictionaryRepository
}

func NewCompanyOptionService() *CompanyOptionService {
	return &CompanyOptionService{repo: repository.NewDictionaryRepository()}
}

func (s *CompanyOptionService) GetCompanyOptions() (*vo.CompanyOptionsVO, error) {
	industries, err := s.repo.GetIndustries()
	if err != nil {
		return nil, err
	}
	natures, err := s.repo.GetNatures()
	if err != nil {
		return nil, err
	}
	scales, err := s.repo.GetScales()
	if err != nil {
		return nil, err
	}

	return &vo.CompanyOptionsVO{
		Industries: industries,
		Natures:    natures,
		Scales:     scales,
	}, nil
}
