package review_errors

import (
	"github.com/MamangRust/monolith-ecommerce-shared/domain/response"

	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcInvalidID = response.NewGrpcError("error", "invalid ID", int(codes.InvalidArgument))

	ErrGrpcValidateCreateReview = response.NewGrpcError("error", "validation failed: invalid create review request", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateReview = response.NewGrpcError("error", "validation failed: invalid update review request", int(codes.InvalidArgument))
)
