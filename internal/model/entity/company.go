package entity

import "time"

type Company struct {
	CompanyID          int       `gorm:"primaryKey;autoIncrement;column:company_id" json:"company_id"`
	UserID             int       `gorm:"unique;not null" json:"user_id"`
	CompanyName        string    `json:"company_name"`
	Description        string    `json:"description"`
	LogoUrl            string    `json:"logo_url"`
	IndustryID         *int      `json:"industry_id"`
	NatureID           *int      `json:"nature_id"`
	CompanyScaleID     *int      `json:"company_scale_id"`
	CompanyAddress     string    `json:"company_address"`
	ContactPersonName  string    `json:"contact_person_name"`
	ContactPersonPhone string    `json:"contact_person_phone"`
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Company) TableName() string {
	return "companies"
}
