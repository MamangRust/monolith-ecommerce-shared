package response_api

import (
	"github.com/MamangRust/monolith-ecommerce-shared/pb"

	"github.com/MamangRust/monolith-ecommerce-shared/domain/response"
)

type merchantResponseMapper struct {
}

func NewMerchantResponseMapper() *merchantResponseMapper {
	return &merchantResponseMapper{}
}

func (m *merchantResponseMapper) ToResponseMerchant(merchant *pb.MerchantResponse) *response.MerchantResponse {
	return &response.MerchantResponse{
		ID:           int(merchant.Id),
		UserID:       int(merchant.UserId),
		Name:         merchant.Name,
		Description:  merchant.Description,
		Address:      merchant.Address,
		ContactEmail: merchant.ContactEmail,
		ContactPhone: merchant.ContactPhone,
		Status:       merchant.Status,
		CreatedAt:    merchant.CreatedAt,
		UpdatedAt:    merchant.UpdatedAt,
	}
}

func (m *merchantResponseMapper) ToResponsesMerchant(merchants []*pb.MerchantResponse) []*response.MerchantResponse {
	var mappedMerchants []*response.MerchantResponse

	for _, merchant := range merchants {
		mappedMerchants = append(mappedMerchants, m.ToResponseMerchant(merchant))
	}

	return mappedMerchants
}

func (m *merchantResponseMapper) ToResponseMerchantDeleteAt(merchant *pb.MerchantResponseDeleteAt) *response.MerchantResponseDeleteAt {
	var deletedAt string
	if merchant.DeletedAt != nil {
		deletedAt = merchant.DeletedAt.Value
	}

	return &response.MerchantResponseDeleteAt{
		ID:           int(merchant.Id),
		UserID:       int(merchant.UserId),
		Name:         merchant.Name,
		Description:  merchant.Description,
		Address:      merchant.Address,
		ContactEmail: merchant.ContactEmail,
		ContactPhone: merchant.ContactPhone,
		Status:       merchant.Status,
		CreatedAt:    merchant.CreatedAt,
		UpdatedAt:    merchant.UpdatedAt,
		DeletedAt:    &deletedAt,
	}
}

func (m *merchantResponseMapper) ToResponsesMerchantDeleteAt(merchants []*pb.MerchantResponseDeleteAt) []*response.MerchantResponseDeleteAt {
	var mappedMerchants []*response.MerchantResponseDeleteAt

	for _, merchant := range merchants {
		mappedMerchants = append(mappedMerchants, m.ToResponseMerchantDeleteAt(merchant))
	}

	return mappedMerchants
}

func (m *merchantResponseMapper) ToApiResponseMerchant(pbResponse *pb.ApiResponseMerchant) *response.ApiResponseMerchant {
	return &response.ApiResponseMerchant{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.ToResponseMerchant(pbResponse.Data),
	}
}

func (m *merchantResponseMapper) ToApiResponseMerchantDeleteAt(pbResponse *pb.ApiResponseMerchantDeleteAt) *response.ApiResponseMerchantDeleteAt {
	return &response.ApiResponseMerchantDeleteAt{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.ToResponseMerchantDeleteAt(pbResponse.Data),
	}
}

func (m *merchantResponseMapper) ToApiResponsesMerchant(pbResponse *pb.ApiResponsesMerchant) *response.ApiResponsesMerchant {
	return &response.ApiResponsesMerchant{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.ToResponsesMerchant(pbResponse.Data),
	}
}

func (m *merchantResponseMapper) ToApiResponseMerchantDelete(pbResponse *pb.ApiResponseMerchantDelete) *response.ApiResponseMerchantDelete {
	return &response.ApiResponseMerchantDelete{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

func (m *merchantResponseMapper) ToApiResponseMerchantAll(pbResponse *pb.ApiResponseMerchantAll) *response.ApiResponseMerchantAll {
	return &response.ApiResponseMerchantAll{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

func (m *merchantResponseMapper) ToApiResponsePaginationMerchantDeleteAt(pbResponse *pb.ApiResponsePaginationMerchantDeleteAt) *response.ApiResponsePaginationMerchantDeleteAt {
	return &response.ApiResponsePaginationMerchantDeleteAt{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       m.ToResponsesMerchantDeleteAt(pbResponse.Data),
		Pagination: *mapPaginationMeta(pbResponse.Pagination),
	}
}

func (m *merchantResponseMapper) ToApiResponsePaginationMerchant(pbResponse *pb.ApiResponsePaginationMerchant) *response.ApiResponsePaginationMerchant {
	return &response.ApiResponsePaginationMerchant{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       m.ToResponsesMerchant(pbResponse.Data),
		Pagination: *mapPaginationMeta(pbResponse.Pagination),
	}
}
