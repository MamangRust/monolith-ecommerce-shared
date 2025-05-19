package response_api

import (
	"github.com/MamangRust/monolith-ecommerce-shared/pb"

	"github.com/MamangRust/monolith-ecommerce-shared/domain/response"
)

type cartResponseMapper struct {
}

func NewCartResponseMapper() *cartResponseMapper {
	return &cartResponseMapper{}
}

func (t *cartResponseMapper) ToResponseCart(pbResponse *pb.CartResponse) *response.CartResponse {
	return &response.CartResponse{
		ID:        int(pbResponse.Id),
		UserID:    int(pbResponse.UserId),
		ProductID: int(pbResponse.ProductId),
		Name:      pbResponse.Name,
		Price:     int(pbResponse.Price),
		Image:     pbResponse.Image,
		Quantity:  int(pbResponse.Quantity),
		Weight:    int(pbResponse.Weight),
		CreatedAt: pbResponse.CreatedAt,
		UpdatedAt: pbResponse.UpdatedAt,
	}
}

func (t *cartResponseMapper) ToResponseCarts(pbResponse []*pb.CartResponse) []*response.CartResponse {
	var carts []*response.CartResponse
	for _, cart := range pbResponse {
		carts = append(carts, t.ToResponseCart(cart))
	}
	return carts
}

func (t *cartResponseMapper) ToApiResponseCart(pbResponse *pb.ApiResponseCart) *response.ApiResponseCart {
	return &response.ApiResponseCart{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    *t.ToResponseCart(pbResponse.Data),
	}
}

func (t *cartResponseMapper) ToApiResponseCartPagination(pbResponse *pb.ApiResponsePaginationCart) *response.ApiResponseCartPagination {
	return &response.ApiResponseCartPagination{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       t.ToResponseCarts(pbResponse.Data),
		Pagination: *mapPaginationMeta(pbResponse.Pagination),
	}
}

func (t *cartResponseMapper) ToApiResponseCartDelete(pbResponse *pb.ApiResponseCartDelete) *response.ApiResponseCartDelete {
	return &response.ApiResponseCartDelete{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

func (t *cartResponseMapper) ToApiResponseCartAll(pbResponse *pb.ApiResponseCartAll) *response.ApiResponseCartAll {
	return &response.ApiResponseCartAll{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}
