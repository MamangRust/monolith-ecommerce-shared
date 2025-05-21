package merchant_errors

import (
	"github.com/MamangRust/monolith-ecommerce-shared/domain/response"

	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcInvalidMerchantId = response.NewGrpcError("error", "invalid merchant ID", int(codes.InvalidArgument))

	ErrGrpcValidateCreateMerchant       = response.NewGrpcError("error", "Validation failed: invalid create merchant request", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateMerchant       = response.NewGrpcError("error", "Validation failed: invalid update merchant request", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateMerchantStatus = response.NewGrpcError("error", "Validation failed: invalid update merchant status request", int(codes.InvalidArgument))
)
