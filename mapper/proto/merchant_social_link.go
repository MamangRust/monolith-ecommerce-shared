package protomapper

import (
	"github.com/MamangRust/monolith-ecommerce-shared/domain/response"
	"github.com/MamangRust/monolith-ecommerce-shared/pb"
)

type merchantSocialLinkProtoMapper struct{}

func NewMerchantSocialLinkProtoMapper() *merchantSocialLinkProtoMapper {
	return &merchantSocialLinkProtoMapper{}
}

func (m *merchantSocialLinkProtoMapper) ToProtoResponseMerchantSocialLink(status string, message string, pbResponse *response.MerchantSocialLinkResponse) *pb.ApiResponseMerchantSocial {
	return &pb.ApiResponseMerchantSocial{
		Status:  status,
		Message: message,
		Data:    m.mapResponseMerchantSocialLink(pbResponse),
	}
}

func (m *merchantSocialLinkProtoMapper) mapResponseMerchantSocialLink(merchant *response.MerchantSocialLinkResponse) *pb.MerchantSocialMediaLinkResponse {
	return &pb.MerchantSocialMediaLinkResponse{
		Id:               int32(merchant.ID),
		MerchantDetailId: int32(merchant.MerchantDetailID),
		Platform:         merchant.Platform,
		Url:              merchant.URL,
		CreatedAt:        merchant.CreatedAt,
		UpdatedAt:        merchant.UpdatedAt,
	}
}
