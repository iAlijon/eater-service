package dtos

type AddressUpdateRequest struct {
	Name      string  `json:"name"`
	AddressID string  `json:"id"`
	Long      float64 `json:"long"`
	Lat       float64 `json:"lat"`
}