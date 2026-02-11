package service

import (
	"job-platform-go/internal/model/vo"
	"job-platform-go/internal/repository"
)

type TagService struct {
	repo *repository.DictionaryRepository
}

func NewTagService() *TagService {
	return &TagService{repo: repository.NewDictionaryRepository()}
}

func (s *TagService) GetAllGroupedTags() (*vo.TagListVO, error) {
	categories, err := s.repo.GetAllTagCategories()
	if err != nil {
		return nil, err
	}
	tags, err := s.repo.GetAllTags()
	if err != nil {
		return nil, err
	}

	// 按 CategoryID 分组
	tagMap := make(map[int][]vo.TagItemVO)
	for _, t := range tags {
		tagMap[t.CategoryID] = append(tagMap[t.CategoryID], vo.TagItemVO{
			TagID:   t.ID,
			TagName: t.Name,
		})
	}

	var groups []vo.CategoryGroupVO
	for _, c := range categories {
		groupTags := tagMap[c.ID]
		if groupTags == nil {
			groupTags = []vo.TagItemVO{}
		}
		groups = append(groups, vo.CategoryGroupVO{
			CategoryID:   c.ID,
			CategoryName: c.Name,
			Tags:         groupTags,
		})
	}

	return &vo.TagListVO{GroupedTags: groups}, nil
}
