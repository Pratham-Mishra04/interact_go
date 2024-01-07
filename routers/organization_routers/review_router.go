package organization_routers

import (
	"github.com/Pratham-Mishra04/interact/controllers/organization_controllers"
	"github.com/Pratham-Mishra04/interact/middlewares"
	"github.com/Pratham-Mishra04/interact/models"
	"github.com/gofiber/fiber/v2"
)

func ReviewRouter(app *fiber.App) {
	app.Get("/org/:orgID/review", organization_controllers.FetchReviews)
	
	reviewRoutes := app.Group("/org/:orgID/review", middlewares.OrgRoleAuthorization(models.Member), middlewares.Protect)
	reviewRoutes.Post("/", organization_controllers.AddReview)
	reviewRoutes.Delete("/", organization_controllers.DeleteReview)
	reviewRoutes.Patch("/like/:reviewID", organization_controllers.LikeReview)
	reviewRoutes.Delete("/like/:reviewID", organization_controllers.RemoveLike)
	reviewRoutes.Patch("/dislike/:reviewID", organization_controllers.DislikeReview)
	reviewRoutes.Delete("/dislike/:reviewID", organization_controllers.RemoveDislike)
}
