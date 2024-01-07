package organization_routers

import (
	"github.com/gofiber/fiber/v2"
)

func Config(app *fiber.App) {
	AuthRouter(app)
	ChatRouter(app)
	EventRouter(app)
	PostRouter(app)
	ProjectRouter(app)
	ProjectApplicationRouter(app)
	ProjectMembershipRouter(app)
	ProjectOpeningRouter(app)
	MembershipRouter(app)
	TaskRouter(app)
	MiscRouter(app)
	ReviewRouter(app)
}
