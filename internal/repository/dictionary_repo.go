package repository

import (
	"gorm.io/gorm"
	"job-platform-go/internal/model/entity"
	"job-platform-go/pkg/database"
)

type DictionaryRepository struct {
	db *gorm.DB
}

func NewDictionaryRepository() *DictionaryRepository {
	return &DictionaryRepository{
		db: database.DB,
	}
}

// FindIndustryIDByName 根据行业名称查找 ID
func (r *DictionaryRepository) FindIndustryIDByName(name string) (int, error) {
	var industry entity.Industry
	err := r.db.Where("name = ?", name).First(&industry).Error
	if err != nil {
		return 0, err
	}
	return industry.ID, nil
}

// GetIndustryNameByID 根据 ID 获取行业名称
func (r *DictionaryRepository) GetIndustryNameByID(id int) (string, error) {
	var industry entity.Industry
	err := r.db.Where("id = ?", id).First(&industry).Error
	if err != nil {
		return "", err
	}
	return industry.Name, nil
}

// FindNatureIDByName 根据企业性质名称查找 ID
func (r *DictionaryRepository) FindNatureIDByName(name string) (int, error) {
	var nature entity.CompanyNature
	err := r.db.Where("name = ?", name).First(&nature).Error
	if err != nil {
		return 0, err
	}
	return nature.ID, nil
}

// GetNatureNameByID 根据 ID 获取性质名称
func (r *DictionaryRepository) GetNatureNameByID(id int) (string, error) {
	var nature entity.CompanyNature
	err := r.db.Where("id = ?", id).First(&nature).Error
	if err != nil {
		return "", err
	}
	return nature.Name, nil
}

// FindScaleIDByName 根据企业规模名称查找 ID
func (r *DictionaryRepository) FindScaleIDByName(name string) (int, error) {
	var scale entity.CompanyScale
	//t_company_scales 表的字段名是 scale 而不是 name
	err := r.db.Where("scale = ?", name).First(&scale).Error
	if err != nil {
		return 0, err
	}
	return scale.ID, nil
}

// GetScaleNameByID 根据 ID 获取规模名称
func (r *DictionaryRepository) GetScaleNameByID(id int) (string, error) {
	var scale entity.CompanyScale
	err := r.db.Where("id = ?", id).First(&scale).Error
	if err != nil {
		return "", err
	}
	return scale.Name, nil
}
