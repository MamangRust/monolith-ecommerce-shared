package response_api

import (
	"github.com/MamangRust/monolith-ecommerce-shared/pb"

	"github.com/MamangRust/monolith-ecommerce-shared/domain/response"
)

type sliderResponseMapper struct {
}

func NewSliderResponseMapper() *sliderResponseMapper {
	return &sliderResponseMapper{}
}

func (s *sliderResponseMapper) ToResponseSlider(pbResponse *pb.SliderResponse) *response.SliderResponse {
	return &response.SliderResponse{
		ID:        int(pbResponse.Id),
		Name:      pbResponse.Name,
		Image:     pbResponse.Image,
		CreatedAt: pbResponse.CreatedAt,
		UpdatedAt: pbResponse.UpdatedAt,
	}
}

func (s *sliderResponseMapper) ToResponsesSlider(pbResponses []*pb.SliderResponse) []*response.SliderResponse {
	var sliders []*response.SliderResponse
	for _, slider := range pbResponses {
		sliders = append(sliders, s.ToResponseSlider(slider))
	}
	return sliders
}

func (s *sliderResponseMapper) ToResponseSliderDeleteAt(pbResponse *pb.SliderResponseDeleteAt) *response.SliderResponseDeleteAt {
	var deletedAt string
	if pbResponse.DeletedAt != nil {
		deletedAt = pbResponse.DeletedAt.Value
	}

	return &response.SliderResponseDeleteAt{
		ID:        int(pbResponse.Id),
		Name:      pbResponse.Name,
		Image:     pbResponse.Image,
		CreatedAt: pbResponse.CreatedAt,
		UpdatedAt: pbResponse.UpdatedAt,
		DeletedAt: &deletedAt,
	}
}

func (s *sliderResponseMapper) ToResponsesSliderDeleteAt(pbResponses []*pb.SliderResponseDeleteAt) []*response.SliderResponseDeleteAt {
	var sliders []*response.SliderResponseDeleteAt
	for _, slider := range pbResponses {
		sliders = append(sliders, s.ToResponseSliderDeleteAt(slider))
	}
	return sliders
}

func (s *sliderResponseMapper) ToApiResponseSlider(pbResponse *pb.ApiResponseSlider) *response.ApiResponseSlider {
	return &response.ApiResponseSlider{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.ToResponseSlider(pbResponse.Data),
	}
}

func (s *sliderResponseMapper) ToApiResponseSliderDeleteAt(pbResponse *pb.ApiResponseSliderDeleteAt) *response.ApiResponseSliderDeleteAt {
	return &response.ApiResponseSliderDeleteAt{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.ToResponseSliderDeleteAt(pbResponse.Data),
	}
}

func (s *sliderResponseMapper) ToApiResponsesSlider(pbResponse *pb.ApiResponsesSlider) *response.ApiResponsesSlider {
	return &response.ApiResponsesSlider{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.ToResponsesSlider(pbResponse.Data),
	}
}

func (s *sliderResponseMapper) ToApiResponseSliderDelete(pbResponse *pb.ApiResponseSliderDelete) *response.ApiResponseSliderDelete {
	return &response.ApiResponseSliderDelete{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

func (s *sliderResponseMapper) ToApiResponseSliderAll(pbResponse *pb.ApiResponseSliderAll) *response.ApiResponseSliderAll {
	return &response.ApiResponseSliderAll{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

func (s *sliderResponseMapper) ToApiResponsePaginationSliderDeleteAt(pbResponse *pb.ApiResponsePaginationSliderDeleteAt) *response.ApiResponsePaginationSliderDeleteAt {
	return &response.ApiResponsePaginationSliderDeleteAt{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       s.ToResponsesSliderDeleteAt(pbResponse.Data),
		Pagination: *mapPaginationMeta(pbResponse.Pagination),
	}
}

func (s *sliderResponseMapper) ToApiResponsePaginationSlider(pbResponse *pb.ApiResponsePaginationSlider) *response.ApiResponsePaginationSlider {
	return &response.ApiResponsePaginationSlider{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       s.ToResponsesSlider(pbResponse.Data),
		Pagination: *mapPaginationMeta(pbResponse.Pagination),
	}
}
