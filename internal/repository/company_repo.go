package repository

import (
	"gorm.io/gorm"
	"job-platform-go/internal/model/entity"
	"job-platform-go/pkg/database"
)

type CompanyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository() *CompanyRepository {
	return &CompanyRepository{
		db: database.DB,
	}
}

func (r *CompanyRepository) FindByUserID(userID int) (*entity.Company, error) {
	var company entity.Company
	err := r.db.Where("user_id=?", userID).First(&company).Error
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *CompanyRepository) GetLinks(companyID int) ([]entity.CompanyExternalLink, error) {
	var companyLinks []entity.CompanyExternalLink
	err := r.db.Where("company_id=?", companyID).First(&companyLinks).Error
	if err != nil {
		return nil, err
	}
	return companyLinks, nil
}

func (r *CompanyRepository) UpdateProfileWithLinks(company *entity.Company, links []entity.CompanyExternalLink) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		//1.更新公司表
		if err := tx.Model(company).
			Where("user_id = ?", company.UserID).
			Updates(map[string]interface{}{
				"description":          company.Description,
				"company_address":      company.CompanyAddress,
				"nature_id":            company.NatureID,
				"industry_id":          company.IndustryID,
				"company_scale_id":     company.CompanyScaleID,
				"contact_person_name":  company.ContactPersonName,
				"contact_person_phone": company.ContactPersonPhone,
			}).Error; err != nil {
			return err
		}

		//2. 删除该企业所有旧链接
		if err := tx.Where("company_id=?", company.CompanyID).Delete(&entity.CompanyExternalLink{}).Error; err != nil {
			return err
		}
		//3. 插入新链接
		if len(links) > 0 {
			for i := range links {
				links[i].ID = 0
				links[i].CompanyID = company.CompanyID
			}
			if err := tx.Create(&links).Error; err != nil {
				return err
			}
		}
		return nil
	})

}
