package dtos

import "github.com/iAlijon/eater-service/src/domain/eater/models"

type EaterSignupResponse struct {
	EaterID string `json:"eater_id"`
}

func NewEaterSignupResponse(eaterID string) *EaterSignupResponse {
	return &EaterSignupResponse{
		EaterID: eaterID,
	}
}