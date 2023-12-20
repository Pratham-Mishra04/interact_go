package routers

import (
	"github.com/Pratham-Mishra04/interact/controllers"
	"github.com/Pratham-Mishra04/interact/controllers/auth_controllers"
	"github.com/Pratham-Mishra04/interact/middlewares"
	"github.com/Pratham-Mishra04/interact/validators"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app *fiber.App) {
	app.Post("/signup", validators.UserCreateValidator, auth_controllers.SignUp)
	app.Post("/login", auth_controllers.LogIn)
	app.Post("/refresh", auth_controllers.Refresh)

	// app.Post("/early_access", auth_controllers.GetEarlyAccessToken)

	app.Post("/recovery", auth_controllers.SendResetURL)
	app.Post("/recovery/verify", auth_controllers.ResetPassword)

	userRoutes := app.Group("/users", middlewares.Protect)
	userRoutes.Get("/me", controllers.GetMe)
	userRoutes.Get("/me/likes", controllers.GetMyLikes)
	userRoutes.Get("/me/organization/memberships", controllers.GetMyOrgMemberships)
	userRoutes.Get("/views", controllers.GetViews)

	userRoutes.Patch("/update_password", controllers.UpdatePassword)
	userRoutes.Patch("/update_email", controllers.UpdateEmail)
	userRoutes.Patch("/update_phone_number", controllers.UpdatePhoneNo)
	userRoutes.Patch("/update_resume", controllers.UpdateResume)

	userRoutes.Get("/get_delete_code", auth_controllers.SendDeleteVerficationCode)
	userRoutes.Delete("/deactive", controllers.Deactive)

	userRoutes.Patch("/me", controllers.UpdateMe)
	userRoutes.Patch("/me/profile", controllers.EditProfile)
	userRoutes.Patch("/me/achievements", controllers.AddAchievement)
	userRoutes.Delete("/me/achievements/:achievementID", controllers.DeleteAchievement)
	// userRoutes.Delete("/me", controllers.DeactivateMe)

	userRoutes.Post("/report", controllers.AddReport)
	userRoutes.Post("/feedback", controllers.AddFeedback)
}
