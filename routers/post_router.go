package routers

import (
	"github.com/Pratham-Mishra04/interact/controllers"
	"github.com/Pratham-Mishra04/interact/middlewares"
	"github.com/gofiber/fiber/v2"
)

func PostRouter(app *fiber.App) {
	postRoutes := app.Group("/posts", middlewares.Protect)
	postRoutes.Post("/", controllers.AddPost)
	postRoutes.Get("/me", controllers.GetMyPosts)
	postRoutes.Get("/me/likes", controllers.GetMyLikedPosts)
	postRoutes.Get("/:postID", controllers.GetPost)
	postRoutes.Patch("/:postID", controllers.UpdatePost)
	postRoutes.Delete("/:postID", controllers.DeletePost)

	postRoutes.Get("/like/:postID", controllers.LikeItem("post"))
	postRoutes.Get("/dislike/:postID", controllers.DislikeItem("post"))
}
