package handler

import (
	"strconv"

	"ddd-boilerplate/pkg/custom_error"
	"ddd-boilerplate/pkg/logger"

	"ddd-boilerplate/pkg/responses"

	"github.com/gofiber/fiber/v2"
)

// GetSampleByID get sample by id
// @Summary Get sample by id
// @Description Get sample data by id
// @Tags Sample
// @Accept  json
// @Param 	id 			path 		int 		true 	"id"
// @Produce json
// @Success 200 {object} responses.ResponseData
// @Success 403 {object} responses.ResponseData
// @Success 500 {object} responses.ResponseData
// @Router /sample/{id} [get]
func (s *Handler) GetSampleByID(ctx *fiber.Ctx) error {
	log := logger.Ctx(ctx.Context())
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		log.Error(err.Error())
		return responses.Error(ctx, custom_error.InvalidIDError)
	}

	sampleResponse, err := s.sampleService.FindSampleByID(ctx.Context(), id)
	if err != nil {
		log.Error(err.Error())
		return responses.Error(ctx, err)
	}

	return ctx.JSON(sampleResponse)
}
