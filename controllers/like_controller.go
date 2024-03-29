package controllers

import (
	"github.com/Pratham-Mishra04/interact/config"
	"github.com/Pratham-Mishra04/interact/helpers"
	"github.com/Pratham-Mishra04/interact/initializers"
	"github.com/Pratham-Mishra04/interact/models"
	"github.com/Pratham-Mishra04/interact/routines"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func handleLikeStatus(likeType string, parsedLoggedInUserID, parsedItemID uuid.UUID, incrementFunc func(uuid.UUID, uuid.UUID), decrementFunc func(uuid.UUID), updateFunc func(uuid.UUID, uuid.UUID)) error {
	var like models.Like

	err := initializers.DB.Where("user_id=? AND "+likeType+"_id=?", parsedLoggedInUserID, parsedItemID).First(&like).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create a new like record
			likeModel := models.Like{
				UserID: parsedLoggedInUserID,
			}
			likeModel.SetItemID(likeType, parsedItemID)

			result := initializers.DB.Create(&likeModel)
			if result.Error != nil {
				return result.Error
			}

			go incrementFunc(parsedItemID, parsedLoggedInUserID)
		} else {
			return err
		}
	} else {
		// Update the like status
		if like.Status == -1 {
			like.Status = 0
			result := initializers.DB.Save(&like)
			if result.Error != nil {
				return result.Error
			}

			go updateFunc(parsedItemID, parsedLoggedInUserID)
		} else {
			// Delete the like record
			result := initializers.DB.Delete(&like)
			if result.Error != nil {
				return result.Error
			}

			go decrementFunc(parsedItemID)
		}
	}

	return nil
}

func handleDislikeStatus(likeType string, parsedLoggedInUserID, parsedItemID uuid.UUID, incrementFunc func(uuid.UUID, uuid.UUID), decrementFunc func(uuid.UUID), updateFunc func(uuid.UUID, uuid.UUID)) error {
	var like models.Like

	err := initializers.DB.Where("user_id=? AND "+likeType+"_id=?", parsedLoggedInUserID, parsedItemID).First(&like).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create a new like record
			likeModel := models.Like{
				UserID: parsedLoggedInUserID,
				Status: -1,
			}
			likeModel.SetItemID(likeType, parsedItemID)

			result := initializers.DB.Create(&likeModel)
			if result.Error != nil {
				return result.Error
			}

			go incrementFunc(parsedItemID, parsedLoggedInUserID)
		} else {
			return err
		}
	} else {
		// Update the like status
		if like.Status == 0 {
			like.Status = -1
			result := initializers.DB.Save(&like)
			if result.Error != nil {
				return result.Error
			}

			go updateFunc(parsedItemID, parsedLoggedInUserID)
		} else {
			// Delete the like record
			result := initializers.DB.Delete(&like)
			if result.Error != nil {
				return result.Error
			}

			go decrementFunc(parsedItemID)
		}
	}

	return nil
}

func LikeItem(likeType string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		loggedInUserID := c.GetRespHeader("loggedInUserID")
		parsedLoggedInUserID, _ := uuid.Parse(loggedInUserID)

		var itemIDParam string
		var incrementFunc func(uuid.UUID, uuid.UUID)
		var decrementFunc func(uuid.UUID)
		var updateFunc func(uuid.UUID, uuid.UUID)

		switch likeType {
		case "post":
			itemIDParam = "postID"
			incrementFunc = routines.IncrementPostLikesAndSendNotification
			decrementFunc = routines.DecrementPostLikes
			updateFunc = func(u1, u2 uuid.UUID) {}

		case "project":
			itemIDParam = "projectID"
			incrementFunc = routines.IncrementProjectLikesAndSendNotification
			decrementFunc = routines.DecrementProjectLikes
			updateFunc = func(u1, u2 uuid.UUID) {}

		case "comment":
			itemIDParam = "commentID"
			incrementFunc = routines.IncrementCommentLikes
			decrementFunc = routines.DecrementCommentLikes
			updateFunc = func(u1, u2 uuid.UUID) {}

		case "event":
			itemIDParam = "eventID"
			incrementFunc = routines.IncrementEventLikesAndSendNotification
			decrementFunc = routines.DecrementEventLikes
			updateFunc = func(u1, u2 uuid.UUID) {}
		
		case "announcement":
			itemIDParam = "announcementID"
			incrementFunc = routines.IncrementAnnouncementLikesAndSendNotification
			decrementFunc = routines.DecrementAnnouncementLikes
			updateFunc = func(u1, u2 uuid.UUID) {}

		case "review":
			itemIDParam = "reviewID"
			incrementFunc = routines.IncrementReviewUpVotes
			decrementFunc = routines.DecrementReviewUpVotes
			updateFunc = func(reviewID, userID uuid.UUID) {
				routines.IncrementReviewUpVotes(reviewID, userID)
				routines.DecrementReviewDownVotes(reviewID)
			}

		default:
			return &fiber.Error{Code: 400, Message: "Invalid likeType"}
		}

		itemID := c.Params(itemIDParam)
		parsedItemID, err := uuid.Parse(itemID)

		if err != nil {
			return &fiber.Error{Code: 400, Message: "Invalid ID"}
		}

		if err := handleLikeStatus(likeType, parsedLoggedInUserID, parsedItemID, incrementFunc, decrementFunc, updateFunc); err != nil {
			return helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
		}

		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Like status updated.",
		})
	}
}

func DislikeItem(likeType string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		loggedInUserID := c.GetRespHeader("loggedInUserID")
		parsedLoggedInUserID, _ := uuid.Parse(loggedInUserID)

		var itemIDParam string
		var incrementFunc func(uuid.UUID, uuid.UUID)
		var decrementFunc func(uuid.UUID)
		var updateFunc func(uuid.UUID, uuid.UUID)

		switch likeType {
		case "post":
			itemIDParam = "postID"
			incrementFunc = func(u1, u2 uuid.UUID) {}
			decrementFunc = func(u1 uuid.UUID) {}
			updateFunc = func(u1, u2 uuid.UUID) {}

		case "project":
			itemIDParam = "projectID"
			incrementFunc = func(u1, u2 uuid.UUID) {}
			decrementFunc = func(u1 uuid.UUID) {}
			updateFunc = func(u1, u2 uuid.UUID) {}

		case "comment":
			itemIDParam = "commentID"
			incrementFunc = func(u1, u2 uuid.UUID) {}
			decrementFunc = func(u1 uuid.UUID) {}
			updateFunc = func(u1, u2 uuid.UUID) {}

		case "event":
			itemIDParam = "eventID"
			incrementFunc = func(u1, u2 uuid.UUID) {}
			decrementFunc = func(u1 uuid.UUID) {}
			updateFunc = func(u1, u2 uuid.UUID) {}

		case "review":
			itemIDParam = "reviewID"
			incrementFunc = routines.IncrementReviewDownVotes
			decrementFunc = routines.DecrementReviewDownVotes
			updateFunc = func(reviewID, userID uuid.UUID) {
				routines.IncrementReviewDownVotes(reviewID, userID)
				routines.DecrementReviewUpVotes(reviewID)
			}

		default:
			return &fiber.Error{Code: 400, Message: "Invalid likeType"}
		}

		itemID := c.Params(itemIDParam)
		parsedItemID, err := uuid.Parse(itemID)

		if err != nil {
			return &fiber.Error{Code: 400, Message: "Invalid ID"}
		}

		if err := handleDislikeStatus(likeType, parsedLoggedInUserID, parsedItemID, incrementFunc, decrementFunc, updateFunc); err != nil {
			return helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
		}

		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Dislike status updated.",
		})
	}
}
