package postgres

import (
	"context"
	"ddd-boilerplate/models"

	"gorm.io/gorm"
)

type SampleRepository struct {
	db *gorm.DB
}

func NewSampleRepository(db *gorm.DB) *SampleRepository {
	return &SampleRepository{
		db: db,
	}
}

func (s *SampleRepository) FindSampleByID(ctx context.Context, id int64) (*models.Sample, error) {
	var sample models.Sample
	result := s.db.WithContext(ctx).Where("id", id).Find(&sample)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &sample, nil
}
