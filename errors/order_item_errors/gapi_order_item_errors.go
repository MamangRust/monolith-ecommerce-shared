package orderitem_errors

import (
	"github.com/MamangRust/monolith-ecommerce-shared/domain/response"

	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcInvalidID = response.NewGrpcError("error", "invalid ID", int(codes.InvalidArgument))
)
