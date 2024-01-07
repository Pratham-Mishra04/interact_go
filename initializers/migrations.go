package initializers

import (
	"fmt"

	"github.com/Pratham-Mishra04/interact/models"
)

func AutoMigrate() {
	fmt.Println("\nStarting Migrations...")
	DB.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Achievement{},
		&models.ProfileView{},
		&models.FollowFollower{},
		&models.College{},

		&models.PostBookmark{},
		&models.PostBookmarkItem{},
		&models.ProjectBookmark{},
		&models.ProjectBookmarkItem{},
		&models.OpeningBookmark{},
		&models.OpeningBookmarkItem{},
		&models.EventBookmark{},
		&models.EventBookmarkItem{},

		&models.Message{},
		&models.GroupChatMessage{},
		&models.Chat{},
		&models.GroupChat{},
		&models.GroupChatMembership{},

		&models.Post{},

		&models.Project{},
		&models.ProjectView{},
		&models.ProjectHistory{},
		&models.Task{},
		&models.SubTask{},
		&models.Opening{},
		&models.Application{},
		&models.Membership{},

		&models.Comment{},
		&models.Invitation{},
		&models.Like{},

		&models.LastViewedProjects{},
		&models.LastViewedOpenings{},

		&models.UserVerification{},
		&models.OAuth{},
		&models.EarlyAccess{},

		&models.Organization{},
		&models.OrganizationMembership{},
		&models.OrganizationHistory{},
		&models.Event{},
		&models.OrganizationReview{},

		&models.Report{},
		&models.Notification{},
		&models.SearchQuery{},
		&models.Feedback{},
	)
	fmt.Println("Migrations Finished!")
}
