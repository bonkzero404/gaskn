package implements

import (
	"gaskn/features/user/factories"
	"gaskn/features/user/repositories"

	"github.com/gofiber/fiber/v2"

	"gaskn/database/stores"
	responseDto "gaskn/dto"
	"gaskn/utils"
)

type UserActivationServiceFactory struct {
	UserActivationRepository repositories.UserActionCodeRepository
}

func NewUserActivationServiceFactory(
	UserActivationRepository repositories.UserActionCodeRepository,
) factories.UserActivationServiceFactory {
	return &UserActivationServiceFactory{
		UserActivationRepository: UserActivationRepository,
	}
}

func (service UserActivationServiceFactory) CreateUserActivation(user *stores.User) (*stores.UserActionCode, error) {
	codeGen := utils.StringWithCharset(32)

	userActivate := stores.UserActionCode{
		UserId:  user.ID,
		Code:    codeGen,
		ActType: stores.ACTIVATION_CODE,
	}

	if err := service.UserActivationRepository.CreateUserActionCode(&userActivate).Error; err != nil {
		return &stores.UserActionCode{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    err.Error(),
		}
	}

	sendMail := responseDto.Mail{
		To:           []string{user.Email},
		Subject:      "User Activation",
		TemplateHtml: "user_activation.html",
		BodyParam: map[string]interface{}{
			"Name": user.FullName,
			"Code": codeGen,
		},
	}

	utils.SendMail(&sendMail)

	return &userActivate, nil
}
