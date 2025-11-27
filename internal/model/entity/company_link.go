package entity

import "time"

type CompanyExternalLink struct {
	ID        int       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	CompanyID int       `gorm:"not null;index" json:"company_id"`
	LinkName  string    `json:"link_name"`
	LinkUrl   string    `json:"link_url"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (CompanyExternalLink) TableName() string {
	return "company_links"
}
