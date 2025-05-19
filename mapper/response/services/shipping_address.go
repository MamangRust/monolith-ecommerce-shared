package response_service

import (
	"github.com/MamangRust/monolith-ecommerce-shared/domain/record"
	"github.com/MamangRust/monolith-ecommerce-shared/domain/response"
)

type shippingAddressResponseMapper struct {
}

func NewShippingAddressResponseMapper() *shippingAddressResponseMapper {
	return &shippingAddressResponseMapper{}
}

func (s *shippingAddressResponseMapper) ToShippingAddressResponse(address *record.ShippingAddressRecord) *response.ShippingAddressResponse {
	return &response.ShippingAddressResponse{
		ID:             address.ID,
		OrderID:        address.OrderID,
		Alamat:         address.Alamat,
		Provinsi:       address.Provinsi,
		Negara:         address.Negara,
		Kota:           address.Kota,
		ShippingMethod: address.ShippingMethod,
		ShippingCost:   int(address.ShippingCost),
		CreatedAt:      address.CreatedAt,
		UpdatedAt:      address.UpdatedAt,
	}
}

func (s *shippingAddressResponseMapper) ToShippingAddressesResponse(addresses []*record.ShippingAddressRecord) []*response.ShippingAddressResponse {
	var responses []*response.ShippingAddressResponse

	for _, address := range addresses {
		responses = append(responses, s.ToShippingAddressResponse(address))
	}

	return responses
}

func (s *shippingAddressResponseMapper) ToShippingAddressResponseDeleteAt(address *record.ShippingAddressRecord) *response.ShippingAddressResponseDeleteAt {
	return &response.ShippingAddressResponseDeleteAt{
		ID:             address.ID,
		OrderID:        address.OrderID,
		Alamat:         address.Alamat,
		Provinsi:       address.Provinsi,
		Negara:         address.Negara,
		Kota:           address.Kota,
		ShippingMethod: address.ShippingMethod,
		ShippingCost:   int(address.ShippingCost),
		CreatedAt:      address.CreatedAt,
		UpdatedAt:      address.UpdatedAt,
		DeletedAt:      address.DeletedAt,
	}
}

func (s *shippingAddressResponseMapper) ToShippingAddressesResponseDeleteAt(addresses []*record.ShippingAddressRecord) []*response.ShippingAddressResponseDeleteAt {
	var responses []*response.ShippingAddressResponseDeleteAt

	for _, address := range addresses {
		responses = append(responses, s.ToShippingAddressResponseDeleteAt(address))
	}

	return responses
}
