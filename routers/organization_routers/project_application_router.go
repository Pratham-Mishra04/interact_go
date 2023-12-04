package organization_routers

import (
	"github.com/Pratham-Mishra04/interact/controllers"
	"github.com/Pratham-Mishra04/interact/middlewares"
	"github.com/Pratham-Mishra04/interact/models"
	"github.com/gofiber/fiber/v2"
)

func ProjectApplicationRouter(app *fiber.App) {
	applicationRoutes := app.Group("/org/:orgID/project/applications", middlewares.OrgProtect, middlewares.OrgRoleAuthorization(models.Manager))

	applicationRoutes.Get("/:applicationID", controllers.GetApplication)

	applicationRoutes.Get("/accept/:applicationID", controllers.AcceptApplication)
	applicationRoutes.Get("/reject/:applicationID", controllers.RejectApplication)
	applicationRoutes.Get("/review/:applicationID", controllers.SetApplicationReviewStatus)

	applicationRoutes.Post("/:openingID", controllers.AddApplication)

	applicationRoutes.Delete("/:applicationID", controllers.DeleteApplication)
}
