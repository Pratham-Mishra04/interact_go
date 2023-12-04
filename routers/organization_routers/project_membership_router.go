package organization_routers

import (
	"github.com/Pratham-Mishra04/interact/controllers"
	"github.com/Pratham-Mishra04/interact/middlewares"
	"github.com/Pratham-Mishra04/interact/models"
	"github.com/gofiber/fiber/v2"
)

func ProjectMembershipRouter(app *fiber.App) {
	membershipRoutes := app.Group("/org/:orgID/project/membership", middlewares.OrgProtect, middlewares.OrgRoleAuthorization(models.Manager))
	membershipRoutes.Post("/:projectID", controllers.AddMember)
	membershipRoutes.Patch("/:membershipID", controllers.ChangeMemberRole)
	membershipRoutes.Delete("/:membershipID", controllers.RemoveMember)
}
