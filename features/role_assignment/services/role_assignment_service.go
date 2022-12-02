package services

import (
	"errors"
	"gaskn/config"
	"gaskn/database/stores"
	respModel "gaskn/dto"
	fclient "gaskn/features/role/contracts"
	"gaskn/features/role_assignment/contracts"
	"gaskn/features/role_assignment/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleAssignmentService struct {
	RoleAssignmentRepository contracts.RoleAssignmentRepository
	RoleClientRepository     fclient.RoleClientRepository
}

func NewRoleAssignmentService(
	RoleAssignmentRepository contracts.RoleAssignmentRepository,
	RoleClientRepository fclient.RoleClientRepository,
) contracts.RoleAssignmentService {
	return &RoleAssignmentService{
		RoleAssignmentRepository: RoleAssignmentRepository,
		RoleClientRepository:     RoleClientRepository,
	}
}

// CreateRoleAssignment /**
func (service RoleAssignmentService) CreateRoleAssignment(c *fiber.Ctx, req *dto.RoleAssignmentRequest) (*dto.RoleAssignmentResponse, error) {
	var clientRole = stores.RoleClient{}

	clientId := c.Params(config.Config("API_CLIENT_PARAM"))

	clientIdUuid, errClientIdUuid := uuid.Parse(clientId)

	roleIdUuid, errRoleUuid := uuid.Parse(req.RoleId)

	if errRoleUuid != nil || errClientIdUuid != nil {
		return &dto.RoleAssignmentResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "invalid format",
		}
	}

	errRoleClient := service.RoleClientRepository.GetRoleClientId(&clientRole, req.RoleId, clientId).Error

	if errors.Is(errRoleClient, gorm.ErrRecordNotFound) {
		return &dto.RoleAssignmentResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    "Role client not found",
		}
	}

	save := service.RoleAssignmentRepository.CreateRoleAssignment(
		roleIdUuid,
		clientIdUuid,
		req.RouteFeature,
		req.MethodFeature,
	)

	if !save {
		return &dto.RoleAssignmentResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "failed to save policy",
		}
	}

	saveResponse := dto.RoleAssignmentResponse{
		RoleId:   roleIdUuid.String(),
		ClientId: clientIdUuid.String(),
	}

	return &saveResponse, nil

}
