package handler

import (
	"ddd-boilerplate/domain/sample"
)

type Handler struct {
	sampleService sample.SampleService
}

func NewHandler(sampleService sample.SampleService) *Handler {
	return &Handler{
		sampleService: sampleService,
	}
}
