package dtos

import "github.com/iAlijon/eater-service/src/domain/address/models"

type GetAddressResponse struct {
	Address *models.Address `json:"address"`
}

func NewGetAddressResponse(address *models.Address) *GetAddressResponse {
	return &GetAddressResponse{
		Address: address,
	}
}