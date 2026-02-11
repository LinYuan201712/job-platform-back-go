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

// === 标签系统 (tag_categories, tags) ===
type TagCategory struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func (TagCategory) TableName() string {
	return "tag_categories"
}

type Tag struct {
	ID         int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string `json:"name"`
	CategoryID int    `json:"category_id"`
	CreatedBy  *int   `json:"created_by"` // 允许自定义标签
}

func (Tag) TableName() string {
	return "tags"
}

// === 地理位置 (t_provinces, t_cities) ===
type Province struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func (Province) TableName() string { return "t_provinces" }

type City struct {
	ID         int    `gorm:"primaryKey;autoIncrement" json:"id"`
	ProvinceID int    `json:"province_id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
}

func (City) TableName() string { return "t_cities" }
