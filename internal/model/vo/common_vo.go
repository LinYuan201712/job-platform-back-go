package vo

type CompanyOptionsVO struct {
	Industries []string `json:"industries"`
	Natures    []string `json:"natures"`
	Scales     []string `json:"scales"`
}

// === Tag VO ===
type TagListVO struct {
	GroupedTags []CategoryGroupVO `json:"grouped_tags"`
}

type CategoryGroupVO struct {
	CategoryID   int         `json:"category_id"`
	CategoryName string      `json:"category_name"`
	Tags         []TagItemVO `json:"tags"`
}

type TagItemVO struct {
	TagID   int    `json:"tag_id"`
	TagName string `json:"tag_name"`
}

// === Location VO ===
type ProvinceVO struct {
	ProvinceID int      `json:"province_id"`
	Name       string   `json:"name"`
	Code       string   `json:"code"`
	Cities     []CityVO `json:"cities"`
}

type CityVO struct {
	CityID int    `json:"city_id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
}
