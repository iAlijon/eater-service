package dtos

import "github.com/iAlijon/eater-service/src/domain/address/models"

type GetAddressListResponse struct {
	Addresses []*models.Address `json:"addresses"`
}

func NewGetAddressListResponse(addresses []*models.Address) *GetAddressListResponse {
	return &GetAddressListResponse{
		Addresses: addresses,
	}
}