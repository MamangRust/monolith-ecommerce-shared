package merchantsociallink_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-ecommerce-shared/domain/response"
	"github.com/labstack/echo/v4"
)

var (
	ErrApiInvalidId = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "invalid merchant social ID", http.StatusBadRequest)
	}

	ErrApiFailedCreateMerchantSocialLink = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "Failed to create merchant social link", http.StatusInternalServerError)
	}

	ErrApiFailedUpdateMerchantSocialLink = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "Failed to update merchant social link", http.StatusInternalServerError)
	}

	ErrApiValidateCreateMerchantSocialLink = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "validation failed: invalid create merchant social link request", http.StatusBadRequest)
	}

	ErrApiValidateUpdateMerchantSocialLink = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "validation failed: invalid update merchant social link request", http.StatusBadRequest)
	}

	ErrApiBindCreateMerchantSocialLink = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "bind failed: invalid create merchant social link request", http.StatusBadRequest)
	}

	ErrApiBindUpdateMerchantSocialLink = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "bind failed: invalid update merchant social link request", http.StatusBadRequest)
	}
)
