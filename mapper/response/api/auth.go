package response_api

import (
	"github.com/MamangRust/monolith-ecommerce-shared/pb"

	"github.com/MamangRust/monolith-ecommerce-shared/domain/response"
)

type authResponseMapper struct {
}

func NewAuthResponseMapper() *authResponseMapper {
	return &authResponseMapper{}
}

func (s *authResponseMapper) ToResponseLogin(res *pb.ApiResponseLogin) *response.ApiResponseLogin {
	if res == nil {
		return &response.ApiResponseLogin{
			Status:  "error",
			Message: "response is nil",
			Data:    nil,
		}
	}

	var tokenResponse *response.TokenResponse
	if res.Data != nil {
		tokenResponse = &response.TokenResponse{
			AccessToken:  res.Data.AccessToken,
			RefreshToken: res.Data.RefreshToken,
		}
	}

	return &response.ApiResponseLogin{
		Status:  res.Status,
		Message: res.Message,
		Data:    tokenResponse,
	}
}

func (s *authResponseMapper) ToResponseRegister(res *pb.ApiResponseRegister) *response.ApiResponseRegister {
	if res == nil {
		return &response.ApiResponseRegister{
			Status:  "error",
			Message: "response is nil",
			Data:    nil,
		}
	}

	var userResponse *response.UserResponse
	if res.Data != nil {
		userResponse = &response.UserResponse{
			ID:        int(res.Data.Id),
			FirstName: res.Data.Firstname,
			LastName:  res.Data.Lastname,
			Email:     res.Data.Email,
			CreatedAt: res.Data.CreatedAt,
			UpdatedAt: res.Data.UpdatedAt,
		}
	}

	return &response.ApiResponseRegister{
		Status:  res.Status,
		Message: res.Message,
		Data:    userResponse,
	}
}

func (s *authResponseMapper) ToResponseRefreshToken(res *pb.ApiResponseRefreshToken) *response.ApiResponseRefreshToken {
	if res == nil {
		return &response.ApiResponseRefreshToken{
			Status:  "error",
			Message: "response is nil",
			Data:    nil,
		}
	}

	var tokenResponse *response.TokenResponse
	if res.Data != nil {
		tokenResponse = &response.TokenResponse{
			AccessToken:  res.Data.AccessToken,
			RefreshToken: res.Data.RefreshToken,
		}
	}

	return &response.ApiResponseRefreshToken{
		Status:  res.Status,
		Message: res.Message,
		Data:    tokenResponse,
	}
}

func (s *authResponseMapper) ToResponseGetMe(res *pb.ApiResponseGetMe) *response.ApiResponseGetMe {
	if res == nil {
		return &response.ApiResponseGetMe{
			Status:  "error",
			Message: "response is nil",
			Data:    nil,
		}
	}

	var userResponse *response.UserResponse
	if res.Data != nil {
		userResponse = &response.UserResponse{
			ID:        int(res.Data.Id),
			FirstName: res.Data.Firstname,
			LastName:  res.Data.Lastname,
			Email:     res.Data.Email,
			CreatedAt: res.Data.CreatedAt,
			UpdatedAt: res.Data.UpdatedAt,
		}
	}

	return &response.ApiResponseGetMe{
		Status:  res.Status,
		Message: res.Message,
		Data:    userResponse,
	}
}
