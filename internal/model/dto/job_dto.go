package dto

import "time"

// 创建新岗位/草稿请求
type HrJobCreateReq struct {
	Title             string `json:"title" binding:"required"`
	Department        string `json:"department"`
	WorkNature        string `json:"work_nature"`
	Type              int    `json:"type"`
	Headcount         int    `json:"headcount"`
	MinSalary         *int   `json:"min_salary"`
	MaxSalary         *int   `json:"max_salary"`
	SalaryRange       string `json:"salary_range"`
	ProvinceID        int    `json:"province_id"`
	CityID            int    `json:"city_id"`
	AddressDetail     string `json:"address_detail"`
	RequiredDegree    int    `json:"required_degree"`
	RequiredStartDate string `json:"required_start_date"`
	Deadline          string `json:"deadline"`
	Description       string `json:"description"`
	TechRequirements  string `json:"tech_requirements"`
	BonusPoints       string `json:"bonus_points"`
	Status            string `json:"status"`
	Tags              []int  `json:"tags"`
	WorkAddress       string `json:"work_address"`
}

// HrJobUpdateReq 更新请求 (字段基本与 Create 一致)
type HrJobUpdateReq struct {
	Title             string `json:"title"`
	Department        string `json:"department"`
	WorkNature        string `json:"work_nature"`
	Type              int    `json:"type"`
	Headcount         int    `json:"headcount"`
	MinSalary         *int   `json:"min_salary"`
	MaxSalary         *int   `json:"max_salary"`
	SalaryRange       string `json:"salary_range"`
	ProvinceID        int    `json:"province_id"`
	CityID            int    `json:"city_id"`
	AddressDetail     string `json:"address_detail"`
	RequiredDegree    int    `json:"required_degree"`
	RequiredStartDate string `json:"required_start_date"`
	Deadline          string `json:"deadline"`
	Description       string `json:"description"`
	TechRequirements  string `json:"tech_requirements"`
	BonusPoints       string `json:"bonus_points"`
	Status            string `json:"status"`
	Tags              []int  `json:"tags"`
	WorkAddress       string `json:"work_address"`
}

// HrJobDetailResp 岗位详情响应
type HrJobDetailResp struct {
	JobID             int         `json:"job_id"`
	Title             string      `json:"title"`
	Status            int         `json:"status"`
	Description       string      `json:"description"`
	TechRequirements  string      `json:"tech_requirements"`
	BonusPoints       string      `json:"bonus_points"`
	MinSalary         *int        `json:"min_salary"`
	MaxSalary         *int        `json:"max_salary"`
	ProvinceID        *int        `json:"province_id"`
	CityID            *int        `json:"city_id"`
	AddressDetail     string      `json:"address_detail"`
	WorkAddress       string      `json:"work_address"`
	WorkNature        *int        `json:"work_nature"`
	Department        string      `json:"department"`
	Headcount         int         `json:"headcount"`
	Type              *int        `json:"type"`
	RequiredDegree    int         `json:"required_degree"`
	RequiredStartDate string      `json:"required_start_date"` // 格式化后的日期字符串
	Deadline          string      `json:"deadline"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	Tags              []JobTagDto `json:"tags"`
}

type JobTagDto struct {
	TagID   int    `json:"tag_id"`
	TagName string `json:"tag_name"`
}

// HrJobListQuery (保持不变)
type HrJobListQuery struct {
	Page         int    `form:"page,default=1"`
	PageSize     int    `form:"page_size,default=10"`
	TitleKeyword string `form:"title_keyword"`
	WorkNature   string `form:"work_nature"`
	Status       string `form:"status"`
}

// HrJobSummary (保持不变)
type HrJobSummary struct {
	JobID       int       `json:"job_id"`
	Title       string    `json:"title"`
	WorkNature  string    `json:"work_nature"`
	Status      string    `json:"status"`
	UpdatedAt   time.Time `json:"updated_at"`
	ReceivedNum int64     `json:"received_num"`
}
