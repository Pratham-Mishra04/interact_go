package organization_routers

import (
	"github.com/Pratham-Mishra04/interact/controllers"
	"github.com/Pratham-Mishra04/interact/controllers/project_controllers"
	"github.com/Pratham-Mishra04/interact/middlewares"
	"github.com/Pratham-Mishra04/interact/models"
	"github.com/gofiber/fiber/v2"
)

func ProjectApplicationRouter(app *fiber.App) {
	applicationRoutes := app.Group("/org/:orgID/project/applications", middlewares.Protect, middlewares.OrgRoleAuthorization(models.Manager))

	applicationRoutes.Get("/:applicationID", controllers.GetApplication("project"))

	applicationRoutes.Get("/accept/:applicationID", project_controllers.AcceptApplication)
	applicationRoutes.Get("/reject/:applicationID", project_controllers.RejectApplication)
	applicationRoutes.Get("/review/:applicationID", project_controllers.SetApplicationReviewStatus)
}
