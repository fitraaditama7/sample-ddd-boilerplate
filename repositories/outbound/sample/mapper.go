package sample

import "ddd-boilerplate/models"

func (s SampleAPIResponse) ToEntity() *models.SampleAPIResponse {
	return &models.SampleAPIResponse{
		ID:                 s.ID,
		Title:              s.Title,
		Description:        s.Description,
		Price:              s.Price,
		DiscountPercentage: s.DiscountPercentage,
		Rating:             s.Rating,
		Stock:              s.Stock,
		Brand:              s.Brand,
		Category:           s.Category,
		Thumbnail:          s.Thumbnail,
		Images:             s.Images,
	}
}
