package handler_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http/httptest"

	"ddd-boilerplate/http/handler"
	"ddd-boilerplate/http/response"
	mocks "ddd-boilerplate/mocks/services"
	"ddd-boilerplate/pkg/custom_error"
	"ddd-boilerplate/pkg/logger"
	"ddd-boilerplate/pkg/responses"
	"github.com/gofiber/fiber/v2"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("SampleHandler", func() {
	var ctrl *handler.Handler
	var sampleService *mocks.SampleService

	BeforeEach(func() {
		sampleService = &mocks.SampleService{}
		ctrl = handler.NewHandler(sampleService)
		logger.InitializeLogger()
	})

	Describe("GetSampleByID", func() {
		table.DescribeTable("success",
			func(expectedId int64, expectedResponse *response.SampleResponse, expectedStatusCode int) {
				sampleService.On("FindSampleByID", mock.Anything, expectedId).Return(expectedResponse, nil).Once()

				app := fiber.New()
				app.Get("/:id", func(c *fiber.Ctx) error {
					return ctrl.GetSampleByID(c)
				})

				req := httptest.NewRequest("GET", fmt.Sprintf("/%d", expectedId), nil)

				resp, err := app.Test(req)
				Expect(err).NotTo(HaveOccurred())

				b, err := io.ReadAll(resp.Body)
				Expect(err).NotTo(HaveOccurred())

				expectedResponseByte, err := json.Marshal(expectedResponse)
				Expect(err).NotTo(HaveOccurred())

				Expect(resp.StatusCode).To(Equal(expectedStatusCode))
				Expect(b).To(Equal(expectedResponseByte))
			},
			table.Entry("unexpected error",
				int64(1),
				&response.SampleResponse{
					ID:   1,
					Name: "test",
				},
				fiber.StatusOK,
			),
		)

		table.DescribeTable("invalid id", func(expectedId string, expectedResponse responses.ResponseData, expectedStatusCode int) {
			sampleService.On("FindSampleByID", mock.Anything, expectedId).Return(nil, nil).Once()
			app := fiber.New()
			app.Get("/:id", func(c *fiber.Ctx) error {
				return ctrl.GetSampleByID(c)
			})

			req := httptest.NewRequest("GET", fmt.Sprintf("/%s", expectedId), nil)

			resp, err := app.Test(req)
			Expect(err).NotTo(HaveOccurred())

			b, err := io.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())

			expectedResponseByte, err := json.Marshal(expectedResponse)
			Expect(err).NotTo(HaveOccurred())

			Expect(resp.StatusCode).To(Equal(expectedStatusCode))
			Expect(string(b)).To(Equal(string(expectedResponseByte)))
		},
			table.Entry("invalid id error",
				"test-invalid-id",
				responses.ResponseData{
					Code:    custom_error.InvalidIDError.ResponseCode,
					Message: custom_error.InvalidIDError.Message,
				},
				fiber.StatusBadRequest,
			))

		table.DescribeTable("service return error",
			func(expectedId int64, expectedResponse responses.ResponseData, expectedStatusCode int, expectedError error) {
				sampleService.On("FindSampleByID", mock.Anything, expectedId).Return(nil, expectedError).Once()

				app := fiber.New()
				app.Get("/:id", func(c *fiber.Ctx) error {
					return ctrl.GetSampleByID(c)
				})

				req := httptest.NewRequest("GET", fmt.Sprintf("/%d", expectedId), nil)

				resp, err := app.Test(req)
				Expect(err).NotTo(HaveOccurred())

				b, err := io.ReadAll(resp.Body)
				Expect(err).NotTo(HaveOccurred())

				expectedResponseByte, err := json.Marshal(expectedResponse)
				Expect(err).NotTo(HaveOccurred())

				Expect(resp.StatusCode).To(Equal(expectedStatusCode))
				Expect(b).To(Equal(expectedResponseByte))
			},
			table.Entry("unexpected error",
				int64(1),
				responses.ResponseData{
					Code:    "Internal Server Error",
					Message: "Internal Server Error",
				},
				fiber.StatusInternalServerError,
				errors.New("unexpected error"),
			),
			table.Entry("bad request error",
				int64(1),
				responses.ResponseData{
					Code:    custom_error.BadRequestError.ResponseCode,
					Message: custom_error.BadRequestError.Message,
				},
				fiber.StatusBadRequest,
				custom_error.BadRequestError,
			),
		)
	})
})
