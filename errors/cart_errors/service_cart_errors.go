package cart_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-ecommerce-shared/domain/response"
)

var (
	ErrCartNotFoundRes   = response.NewErrorResponse("Cart not found", http.StatusNotFound)
	ErrCartAlreadyExists = response.NewErrorResponse("Cart already exists", http.StatusBadRequest)
	ErrCartInvalidData   = response.NewErrorResponse("Invalid cart data", http.StatusBadRequest)

	ErrFailedFindAllCarts = response.NewErrorResponse("Failed to fetch carts", http.StatusInternalServerError)
	ErrFailedCreateCart   = response.NewErrorResponse("Failed to create cart", http.StatusInternalServerError)

	ErrFailedDeleteCart     = response.NewErrorResponse("Failed to permanently delete cart", http.StatusInternalServerError)
	ErrFailedDeleteAllCarts = response.NewErrorResponse("Failed to permanently delete all carts", http.StatusInternalServerError)
)
