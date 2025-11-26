package entity

import "time"

type Student struct {
	UserID            int        `gorm:"primaryKey;column:user_id" json:"user_id"` // 手动指定主键为 user_id
	StudentID         string     `gorm:"unique;column:student_id" json:"student_id"`
	AvatarUrl         string     `json:"avatar_url"`
	FullName          string     `json:"full_name"`
	PhoneNumber       string     `json:"phone_number"`
	Gender            *int       `json:"gender"` // 0=男, 1=女
	DateOfBirth       *time.Time `gorm:"type:date" json:"date_of_birth"`
	JobSeekingStatus  *int       `json:"job_seeking_status"`
	ExpectedPosition  string     `json:"expected_position"`
	ExpectedMinSalary *int       `json:"expected_min_salary"`
	ExpectedMaxSalary *int       `json:"expected_max_salary"`
	SkillsSummary     string     `json:"skills_summary"`
	CurrentTemplateID *int64     `json:"current_template_id"`
}

func (Student) TableName() string {
	return "students"
}
