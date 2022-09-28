package role

import (
	"go-starterkit-project/app/middleware"
	"go-starterkit-project/modules/role/handlers"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
)

type ApiRoute struct {
	RoleHandler       handlers.RoleHandler
	RoleClientHandler handlers.RoleClientHandler
}

func (handler *ApiRoute) Route(app fiber.Router) {
	const endpointGroup string = "/role"

	role := app.Group(utils.SetupApiGroup() + endpointGroup)
	roleClient := app.Group(utils.SetupSubApiGroup() + endpointGroup)

	///////////////////
	// Route Role
	///////////////////
	role.Post(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleHandler.CreateRole,
	).Name("CreateRole")

	role.Get(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleHandler.GetRoleList,
	).Name("GetRoleList")

	role.Put(
		"/:id",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleHandler.UpdateRole,
	).Name("UpdateRole")

	role.Delete(
		"/:id",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleHandler.DeleteRole,
	).Name("DeletRole")

	//////////////////////
	// Route Role Client
	//////////////////////
	roleClient.Post(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleClientHandler.CreateClientRole,
	).Name("CreateClientRole")

	roleClient.Get(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleClientHandler.GetRoleClientList,
	).Name("GetRoleClientList")

	roleClient.Put(
		"/:id",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleClientHandler.UpdateRoleClient,
	).Name("UpdateRoleCLient")

	roleClient.Delete(
		"/:id",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleClientHandler.DeleteRoleClient,
	).Name("DeleteRoleClient")

}
