package repositories

import (
	"go-starterkit-project/domain/stores"
	"go-starterkit-project/modules/role/domain/interfaces"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) interfaces.RoleRepositoryInterface {
	return &RoleRepository{
		DB: db,
	}
}

func (repository RoleRepository) CreateRole(role *stores.Role) *gorm.DB {
	return repository.DB.Create(&role)
}

func (repository RoleRepository) UpdateRoleById(role *stores.Role) *gorm.DB {
	return repository.DB.Save(&role)
}

func (repository RoleRepository) DeleteRoleById(role *stores.Role) *gorm.DB {
	return repository.DB.Delete(&role)
}

func (repository RoleRepository) GetRoleById(role *stores.Role, id string) *gorm.DB {
	return repository.DB.First(&role, "id = ?", id)
}

func (repository RoleRepository) GetRoleList(role *[]stores.Role, c *fiber.Ctx) (*utils.Pagination, error) {
	var pagination utils.Pagination

	err := repository.DB.Scopes(utils.Paginate(role, &pagination, repository.DB, c)).Find(&role).Error

	return &pagination, err
}

func (repository RoleRepository) GetClientsByUser(roleClient *[]stores.RoleUser, userId string) *gorm.DB {
	return repository.DB.
		Select("client_id", "role_id").
		Where("user_id = ?", userId).
		Group("client_id, role_id").
		Preload("Client").
		Preload("Role").
		Find(&roleClient)
}
