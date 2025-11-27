package dto

import "time"

type ExternalLinkDto struct {
	LinkName string `json:"link_name" binding:"required"`
	LinkUrl  string `json:"link_url" binding:"required"`
}

// 更新企业信息请求
type CompanyProfileUpdateReq struct {
	Description        string            `json:"description" binding:"required"`
	CompanyAddress     string            `json:"company_address" binding:"required"`
	Nature             string            `json:"nature" binding:"required"`        // 例如：私营企业
	Industry           string            `json:"industry" binding:"required"`      // 例如：互联网
	CompanyScale       string            `json:"company_scale" binding:"required"` // 例如：100-499人
	ContactPersonName  string            `json:"contact_person_name" binding:"required"`
	ContactPersonPhone string            `json:"contact_person_phone" binding:"required"`
	ExternalLinks      []ExternalLinkDto `json:"external_links"`
}

// 企业信息回显响应
type CompanyProfileResp struct {
	CompanyName        string `json:"company_name"`
	Description        string `json:"description"`
	LogoUrl            string `json:"logo_url"`
	Nature             string `json:"nature"`
	Industry           string `json:"industry"`
	CompanyScale       string `json:"company_scale"`
	ContactPersonName  string `json:"contact_person_name"`
	ContactPersonPhone string `json:"contact_person_phone"`
	CompanyAddress     string `json:"company_address"`

	// 统计数据 (暂时 Mock，后续对接 Job/Application 模块)
	OpenJobsCount     int64      `json:"open_jobs_count"`
	ResumeProcessRate string     `json:"resume_process_rate"`
	LastLoginAt       *time.Time `json:"last_login_at"`

	ExternalLinks []ExternalLinkDto `json:"external_links"`
}
