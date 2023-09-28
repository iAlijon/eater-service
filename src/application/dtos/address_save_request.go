package dtos

type AddressSaveRequest struct {
	Name    string  `json:"name"`
	EaterID string  `json:"eater_id"`
	Long    float64 `json:"long"`
	Lat     float64 `json:"lat"`
}