package entity

// === 企业字典 ===
// 行业领域 (如：互联网、金融)
type Industry struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

func (Industry) TableName() string {
	return "t_industries"
}

// 企业性质 (如：国企、外企)
type CompanyNature struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

func (CompanyNature) TableName() string {
	return "t_company_natures"
}

// 人员规模 (如：0-20人、10000人以上)
type CompanyScale struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

func (CompanyScale) TableName() string {
	return "t_company_scales"
}
