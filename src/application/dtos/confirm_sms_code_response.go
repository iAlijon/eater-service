package dtos

import "github.com/iAlijon/eater-service/src/domain/eater/models"
type ConfirmSMSCodeResponse struct {
	Token string `json:"token"`
	Profile *models.EaterProfile `json:"profile"`
}