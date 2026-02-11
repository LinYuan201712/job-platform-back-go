package service

import (
	"job-platform-go/internal/model/vo"
	"job-platform-go/internal/repository"
)

type LocationService struct {
	repo *repository.DictionaryRepository
}

func NewLocationService() *LocationService {
	return &LocationService{repo: repository.NewDictionaryRepository()}
}

func (s *LocationService) GetProvincesCities() ([]vo.ProvinceVO, error) {
	// 并行查询或者顺序查询均可，这里用顺序查询
	provinces, err := s.repo.GetAllProvinces()
	if err != nil {
		return nil, err
	}
	cities, err := s.repo.GetAllCities()
	if err != nil {
		return nil, err
	}

	// 内存中组装树形结构，比 N+1 次 SQL 查询效率高得多
	cityMap := make(map[int][]vo.CityVO)
	for _, c := range cities {
		cityMap[c.ProvinceID] = append(cityMap[c.ProvinceID], vo.CityVO{
			CityID: c.ID,
			Name:   c.Name,
			Code:   c.Code,
		})
	}

	var result []vo.ProvinceVO
	for _, p := range provinces {
		cities := cityMap[p.ID]
		if cities == nil {
			cities = []vo.CityVO{} // 避免前端拿到 null
		}
		result = append(result, vo.ProvinceVO{
			ProvinceID: p.ID,
			Name:       p.Name,
			Code:       p.Code,
			Cities:     cities,
		})
	}
	return result, nil
}
