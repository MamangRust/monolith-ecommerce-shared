package response_service

import (
	"github.com/MamangRust/monolith-ecommerce-shared/domain/record"
	"github.com/MamangRust/monolith-ecommerce-shared/domain/response"
)

type merchantSocialLinkResponseMapper struct {
}

func NewMerchantSocialLinkResponseMapper() *merchantSocialLinkResponseMapper {
	return &merchantSocialLinkResponseMapper{}
}

func (s *merchantSocialLinkResponseMapper) ToMerchantSocialLinkResponse(merchant *record.MerchantSocialLinkRecord) *response.MerchantSocialLinkResponse {
	return &response.MerchantSocialLinkResponse{
		ID:               merchant.ID,
		MerchantDetailID: merchant.MerchantDetailID,
		Platform:         merchant.Platform,
		URL:              merchant.URL,
		CreatedAt:        merchant.CreatedAt,
		UpdatedAt:        merchant.UpdatedAt,
	}
}
