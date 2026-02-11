package repository

import (
	"gorm.io/gorm"
	"job-platform-go/internal/model/entity"
	"job-platform-go/pkg/database"
)

type JobRepository struct {
	db *gorm.DB
}

func NewJobRepository() *JobRepository {
	return &JobRepository{
		db: database.DB,
	}
}
func (r *JobRepository) CreateJob(job *entity.Job, tagIDs []int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(job).Error; err != nil {
			return err
		}
		return r.replaceTags(tx, job, tagIDs)
	})
}
func (r *JobRepository) replaceTags(tx *gorm.DB, job *entity.Job, tagIDs []int) error {
	if len(tagIDs) > 0 {
		tags := make([]entity.Tag, len(tagIDs))
		for i, id := range tagIDs {
			tags[i] = entity.Tag{ID: id}
		}
		// GORM Association 自动处理 job_tags 中间表的增删
		return tx.Model(job).Association("Tags").Replace(tags)
	}
	// 如果 tagIDs 为空数组，清空关联
	return tx.Model(job).Association("Tags").Clear()
}

// FindByIDAndCompany 查找岗位（带权限校验）
func (r *JobRepository) FindByIDAndCompany(jobID, companyID int) (*entity.Job, error) {
	var job entity.Job
	// Preload("Tags") 自动查询关联表
	err := r.db.Preload("Tags").Where("id = ? AND company_id = ?", jobID, companyID).First(&job).Error
	return &job, err
}
func (r *JobRepository) Update(job *entity.Job, tagIDs []int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Save 会更新所有字段，包括零值
		if err := tx.Save(job).Error; err != nil {
			return err
		}
		if tagIDs != nil {
			return r.replaceTags(tx, job, tagIDs)
		}
		return nil
	})
}
