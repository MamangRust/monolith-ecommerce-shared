package merchantsociallink_errors

import (
	"github.com/MamangRust/monolith-ecommerce-shared/domain/response"
	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcMerchantSocialLinkInvalidID = response.NewGrpcError(
		"merchant_social_link",
		"Invalid merchant social link id",
		int(codes.InvalidArgument),
	)

	ErrGrpcFailedCreateMerchantSocialLink = response.NewGrpcError(
		"merchant_social_link",
		"Failed to create merchant social link",
		int(codes.Internal),
	)

	ErrGrpcFailedUpdateMerchantSocialLink = response.NewGrpcError(
		"merchant_social_link",
		"Failed to update merchant social link",
		int(codes.Internal),
	)

	ErrGrpcValidateCreateMerchantSocialLink = response.NewGrpcError(
		"merchant_social_link",
		"Invalid input for create merchant social link",
		int(codes.InvalidArgument),
	)

	ErrGrpcValidateUpdateMerchantSocialLink = response.NewGrpcError(
		"merchant_social_link",
		"Invalid input for update merchant social link",
		int(codes.InvalidArgument),
	)
)
