package services

import (
	"context"
	"fmt"
	"time"
	"github.com/iAlijon/eater-service/src/domain/eater/models"
	"github.com/iAlijon/eater-service/src/domain/eater/repositories"
	"go.uber.org/zap"
)

type EaterService interface {
	SignupEater(ctx context.Context, phoneNumber string) (string, error)
	ConfirmSMSCode(ctx context.Context, eaterID, smsCode string) (*models.EaterProfile, error)
	GetEaterProfile(ctx context.Context, eaterID string) (*models.EaterProfile, error)
	UpdateEaterProfile(ctx context.Context, eaterID, name, imageUrl string) (*models.EaterProfile, error)
}

type eaterSvcImpl struct {
	eaterRepo repositories.EaterRepository
	smsClient sms.Client
	logger *zap.Logger
}

func NewEaterService(
	eaterRepo repositories.EaterRepository,
	smsClient sms.Client,
	logger *zap.Logger,
) EaterService {
	return &eaterSvcImpl{
		eaterRepo: eaterRepo,
		smsClient: smsClient,
		logger: logger,
	}
}

func (s *eaterSvcImpl) SignupEater(ctx context.Context, phoneNumber sting) (string, error)  {
	phoneNumberExist := true

	eater, err := s.eaterRepo.GetEaterByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		phoneNumberExist = false
	}
	if phoneNumberExist{
		return s.handleExistingEater(ctx, eater.ID)
	}
	return s.handleNewEater(ctx, phoneNumber)

}

func (s *eaterSvcImpl) handleNewEater(ctx contect.Context, phoneNumber string) (string, error){
	var (
		eaterID = rand.UUID()
		earterName = fmt.Sprintf("eater-%s", rand.NumericString(5))
		salt = crypto.GenerateSalt()
		saltedPass = crypto.Combine(salt, phoneNumber)
		passHash = crypto.HashPassword(saltedPass)
		now = time.Now().UTC()
	)
	eater := models.Eater{
		ID: eaterID,
		PhoneNumeber: phoneNumber,
		PasswordHash: passHash,
		PasswordSalt: salt,
		CreatedAt: now,
		UpdateAt: now,
	}

	eaterProfile := models.EaterProfile{
		EaterID: eaterID,
		PhoneNumber: phoneNumber,
		Name: earterName,
		ImageUrl: "",
		CreatedAt: now,
		UpdateAt: now,
	}

	smsCode := models.EaterSmsCode{
		EaterID: eaterID,
		Code: rand.NumericString(5),
		ExpiresIn: 300,
		CreatedAt: now,
	}

	err := s.eaterRepo.WithTx(ctx, func(r repositories.EaterRepository) error {
		if err := r.SaveEater(ctx, &eater); err != nil {
			return err
		}
		if err := r.SaveEaterProfile(ctx, &eaterProfile); err != nil {
			return err
		}
		if err := r.SaveEaterSmsCode(ctx, &smsCode); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	smsMsg := fmt.Sprintf("Food.uz Code: %s", smsCode.Code)
	if err := s.smsClient.SandMessage(ctx, eater.phoneNumber,smsMsg); err != nil {
		return "",err
	}
	return eaterID,nil

}

func (s *eaterSvcImpl) handleExistingEater(ctx context.Context, eaterID string) (string, error)  {
	eater, err := s.eaterRepo.GetEater(ctx,eaterID)
	if err != nil {
		return "", nil
	}

	smsCode := models.EaterSmsCode{
		EaterID: eaterID,
		Code: rand.NumericString(5),
		ExpiresIn: 300,
		CreatedAt: now,
	}
	if err := s.eaterRepo.SaveEaterSmsCode(ctx, &smsCode); err != nil {
		return "", err
	}
	smsMsg := fmt.Srpintf("Food.uz Code: %s", smsCode.Code)
	if err := s.smsClient.SendMassage(ctx, eater.PhoneNumber, smsMsg); err != nil {
		return "",err
	}
	return eaterID, nil
}

func (s *eaterSvcImpl) ConfirSMSCode(ctx context.Context, eaterID , smsCode string) (string, error)  {
	smsCode, err := s.eaterRepo.GetEaterSmsCode(ctx, eaterID, code)
	if err !=nil {
		return nil, err
	}
	return eaterID, nil
}

func (s *eaterSvcImpl) GetEaterProfile(ctx context.Context, eaterID string) (string, error)  {
	return nil, nil
}

func (s *eaterSvcImpl)UpdateEaterProfile(ctx context.Context, eaterID, name, imageUrl string) (string, error)  {
	return nil, nil
}





