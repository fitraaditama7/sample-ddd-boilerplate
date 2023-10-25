package sample

import (
	"context"
	"ddd-boilerplate/repositories/outbound/sample"
)

//go:generate mockery --dir=./ --name=SampleOutboundRepository --filename=sample_outbound_repository.go --output=../../mocks/repository --outpkg=mock_repository
type SampleOutboundRepository interface {
	FindSampleAPI(ctx context.Context, id int64) (*sample.SampleAPIResponse, error)
	PostSampleAPI(ctx context.Context, request sample.SamplePostAPIRequest) (*sample.SamplePostAPIResponse, error)
}
