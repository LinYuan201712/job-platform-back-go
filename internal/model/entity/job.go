package entity

import (
	"time"
)

type Job struct {
	ID                int        `gorm:"primaryKey;autoIncrement" json:"id"`
	CompanyID         int        `gorm:"not null;index" json:"company_id"`        // 关联 Company
	PostedByUserID    int        `gorm:"not null;index" json:"posted_by_user_id"` // 发布人的 UserID
	Title             string     `gorm:"size:255" json:"title"`
	Description       string     `gorm:"type:text" json:"description"`
	TechRequirements  string     `gorm:"type:text" json:"tech_requirements"`
	MinSalary         *int       `json:"min_salary"` // 薪资单位：k
	MaxSalary         *int       `json:"max_salary"`
	ProvinceID        *int       `json:"province_id"`
	CityID            *int       `json:"city_id"`
	AddressDetail     string     `json:"address_detail"`
	WorkNature        *int       `json:"work_nature"` // 1=实习, 2=校招
	Headcount         int        `json:"headcount"`
	Status            int        `gorm:"default:10" json:"status"` // 默认状态
	Type              *int       `json:"type"`                     // 职能类别 ID (关联 JobCategory)
	Department        string     `json:"department"`
	RequiredDegree    int        `json:"required_degree"` // 0=本科, 1=硕士...
	RequiredStartDate *time.Time `gorm:"type:date" json:"required_start_date"`
	BonusPoints       string     `gorm:"type:text" json:"bonus_points"` // 加分项，存 JSON 或文本
	WorkAddress       string     `json:"work_address"`
	ViewCount         int        `gorm:"default:0" json:"view_count"` // 浏览量

	Deadline  *time.Time `gorm:"type:date" json:"deadline"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`

	// --- 关联关系 ---

	// 关联企业 (Belongs To)
	Company Company `gorm:"foreignKey:CompanyID;references:CompanyID" json:"company,omitempty"`

	// 关联标签 (Many To Many)
	// GORM 会自动管理中间表 `job_tags`
	Tags []Tag `gorm:"many2many:job_tags;joinForeignKey:JobID;joinReferences:TagID" json:"tags,omitempty"`
}

// JobStatus岗位状态枚举
const (
	JobStatusDraft    = 1  // 草稿
	JobStatusPending  = 10 // 待审核
	JobStatusApproved = 20 // 已发布
	JobStatusRejected = 30 // 已驳回
	JobStatusClosed   = 40 // 已关闭/下线
)

// WorkNature工作性质枚举
const (
	WorkNatureInternship = 1 // 实习
	WorkNatureFullTime   = 2 // 校招/全职
)
