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

// === Company Options ===
// 使用 Pluck 只查某一列，减少内存消耗
func (r *DictionaryRepository) GetIndustries() ([]string, error) {
	var names []string
	err := database.DB.Model(&entity.Industry{}).Order("id asc").Pluck("name", &names).Error
	return names, err
}

func (r *DictionaryRepository) GetNatures() ([]string, error) {
	var names []string
	err := database.DB.Model(&entity.CompanyNature{}).Order("id asc").Pluck("name", &names).Error
	return names, err
}

func (r *DictionaryRepository) GetScales() ([]string, error) {
	var scales []string
	err := database.DB.Model(&entity.CompanyScale{}).Order("id asc").Pluck("scale", &scales).Error
	return scales, err
}

// === Tags ===
// GetAllTagCategories获取所有的标签分类
func (r *DictionaryRepository) GetAllTagCategories() ([]entity.TagCategory, error) {
	var categories []entity.TagCategory
	err := database.DB.Order("id asc").Find(&categories).Error
	return categories, err
}

// GetAllTags获取所有的标签
func (r *DictionaryRepository) GetAllTags() ([]entity.Tag, error) {
	var tags []entity.Tag
	err := database.DB.Order("id asc").Find(&tags).Error
	return tags, err
}

// === Location ===
func (r *DictionaryRepository) GetAllProvinces() ([]entity.Province, error) {
	var provinces []entity.Province
	err := database.DB.Order("id asc").Find(&provinces).Error
	return provinces, err
}

func (r *DictionaryRepository) GetAllCities() ([]entity.City, error) {
	var cities []entity.City
	err := database.DB.Order("id asc").Find(&cities).Error
	return cities, err
}
