package sample_test

import (
	"context"
	"ddd-boilerplate/domain/sample"
	"errors"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/ginkgo/v2/dsl/table"
	. "github.com/onsi/gomega"

	"ddd-boilerplate/http/response"
	mockRepo "ddd-boilerplate/mocks/repository"
	models "ddd-boilerplate/models"
	"ddd-boilerplate/pkg/logger"
)

var _ = Describe("sample test", func() {
	var (
		sampleRepo    *mockRepo.SampleOutboundRepository
		sampleService sample.SampleService
	)

	BeforeEach(func() {
		sampleRepo = mockRepo.NewSampleOutboundRepository(GinkgoT())
		sampleService = sample.NewSampleService(sampleRepo)
		logger.InitializeLogger()
	})

	Describe("find sample by id", func() {
		ctx := context.Background()
		table.DescribeTable("found the sample", func(expectedID int64, expectedReturn *response.SampleResponse) {
			sampleRepo.On("FindSampleAPI", ctx, int64(1)).Return(&models.SampleAPIResponse{
				ID:    1,
				Brand: "ahah",
			}, nil)

			res, err := sampleService.FindSampleByID(ctx, expectedID)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(expectedReturn))
		},
			table.Entry("success", int64(1), &response.SampleResponse{
				ID:   1,
				Name: "ahah",
			}),
		)

		table.DescribeTable("error", func(expectedID int64, expectedReturn *response.SampleResponse) {
			sampleRepo.On("FindSampleAPI", ctx, int64(1)).Return(nil, errors.New("error"))

			_, err := sampleService.FindSampleByID(ctx, expectedID)
			Expect(err).To(Equal(errors.New("internal server error")))
		},
			table.Entry("success", int64(1), nil),
		)
	})

})
