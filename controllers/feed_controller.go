package controllers

import (
	"github.com/Pratham-Mishra04/interact/config"
	"github.com/Pratham-Mishra04/interact/helpers"
	"github.com/Pratham-Mishra04/interact/initializers"
	"github.com/Pratham-Mishra04/interact/models"
	"github.com/Pratham-Mishra04/interact/routines"
	API "github.com/Pratham-Mishra04/interact/utils/APIFeatures"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetFeed(c *fiber.Ctx) error {
	loggedInUserID := c.GetRespHeader("loggedInUserID")

	var followings []models.FollowFollower
	if err := initializers.DB.Model(&models.FollowFollower{}).Where("follower_id = ?", loggedInUserID).Find(&followings).Error; err != nil {
		return helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	followingIDs := make([]uuid.UUID, len(followings))
	for i, following := range followings {
		followingIDs[i] = following.FollowedID
	}

	paginatedDB := API.Paginator(c)(initializers.DB)

	var posts []models.Post
	if err := paginatedDB.
		Preload("User").
		Preload("RePost").
		Preload("RePost.User").
		Preload("RePost.TaggedUsers").
		Preload("TaggedUsers").
		Joins("JOIN users ON posts.user_id = users.id AND users.active = ?", true).
		Where("user_id = ? OR user_id IN (?)", loggedInUserID, followingIDs).
		Order("created_at DESC").
		Find(&posts).Error; err != nil {
		return helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	go routines.IncrementPostImpression(posts)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"feed":   posts,
	})
}
