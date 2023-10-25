package sample

import (
	"context"
	models "ddd-boilerplate/models"
	"ddd-boilerplate/pkg/logger"
	sample_outbound "ddd-boilerplate/repositories/outbound/sample"
	"errors"
)

type sampleService struct {
	// sampleRepository         repository.SampleRepository
	sampleOutboundRepository SampleOutboundRepository
}

//go:generate mockery --name SampleService --output ../../mocks/services
type SampleService interface {
	FindSampleByID(ctx context.Context, id int64) (*models.SampleAPIResponse, error)
}

func NewSampleService(outboundRepository SampleOutboundRepository) SampleService {
	return &sampleService{
		sampleOutboundRepository: outboundRepository,
	}
}

func (s *sampleService) FindSampleByID(ctx context.Context, id int64) (*models.SampleAPIResponse, error) {
	log := logger.Ctx(ctx)
	sample, err := s.sampleOutboundRepository.FindSampleAPI(ctx, id)
	if err != nil {
		log.Error(err.Error())
		return nil, errors.New("internal server error")
	}

	_, err = s.sampleOutboundRepository.PostSampleAPI(ctx, sample_outbound.SamplePostAPIRequest{Title: "test title"})
	if err != nil {
		log.Error(err.Error())
		return nil, errors.New("internal server error")
	}

	result := sample.ToEntity()
	return result, nil
}
