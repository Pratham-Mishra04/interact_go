package organization_controllers

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

func AcceptApplication(c *fiber.Ctx) error {
	applicationID := c.Params("applicationID")
	loggedInUserID := c.GetRespHeader("loggedInUserID")
	orgMemberID := c.GetRespHeader("orgMemberID")
	parsedOrgMemberID, _ := uuid.Parse(orgMemberID)

	parsedApplicationID, err := uuid.Parse(applicationID)
	if err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid ID"}
	}

	var application models.Application
	if err := initializers.DB.Preload("Opening").Preload("Opening.Organization").First(&application, "id = ?", parsedApplicationID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &fiber.Error{Code: 400, Message: "No Application of this ID found."}
		}
		return helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	if application.Opening.UserID.String() != loggedInUserID {
		return &fiber.Error{Code: 403, Message: "Do not have the permission to perform this action."}
	}

	if application.Status == -1 {
		return &fiber.Error{Code: 400, Message: "Application is already Rejected."}
	}

	application.Status = 2
	result := initializers.DB.Save(&application)

	if result.Error != nil {
		return helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	membership := models.OrganizationMembership{
		OrganizationID: *application.Opening.OrganizationID,
		UserID:    application.UserID,
		Role:      models.Member,
		Title:     application.Opening.Title,
	}

	result = initializers.DB.Create(&membership)

	if result.Error != nil {
		helpers.LogDatabaseError("Error while creating Membership-CreateOrgMembershipAndSendNotification", result.Error, "go_routine")
	}

	go routines.OrgMembershipSendNotification(&application)

	go routines.MarkOrganizationHistory(*application.ProjectID, parsedOrgMemberID, 27, &application.UserID, nil, nil, nil, nil, nil, nil, &application.OpeningID, "")

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Application Accepted.",
	})
}

func RejectApplication(c *fiber.Ctx) error {
	applicationID := c.Params("applicationID")
	loggedInUserID := c.GetRespHeader("loggedInUserID")
	parsedLoggedInUserID, _ := uuid.Parse(loggedInUserID)
	orgMemberID := c.GetRespHeader("orgMemberID")
	parsedOrgMemberID, _ := uuid.Parse(orgMemberID)

	parsedApplicationID, err := uuid.Parse(applicationID)
	if err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid ID"}
	}

	var application models.Application
	if err := initializers.DB.Preload("Opening").First(&application, "id = ?", parsedApplicationID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &fiber.Error{Code: 400, Message: "No Application of this ID found."}
		}
		return helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	if application.Opening.UserID.String() != loggedInUserID {
		return &fiber.Error{Code: 403, Message: "Do not have the permission to perform this action."}
	}

	if application.Status == 2 {
		return &fiber.Error{Code: 400, Message: "Application is already Accepted."}
	}

	application.Status = -1
	result := initializers.DB.Save(&application)

	if result.Error != nil {
		return helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	notification := models.Notification{
		NotificationType: 7,
		UserID:           application.UserID,
		SenderID:         parsedLoggedInUserID,
		OpeningID:        &application.OpeningID,
	}

	if err := initializers.DB.Create(&notification).Error; err != nil {
		return helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	go routines.MarkOrganizationHistory(*application.OrganizationID, parsedOrgMemberID, 28, nil, nil, nil, nil, nil, nil, nil, nil, application.Opening.Title)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Application Rejected.",
	})
}

func SetApplicationReviewStatus(c *fiber.Ctx) error {
	applicationID := c.Params("applicationID")
	loggedInUserID := c.GetRespHeader("loggedInUserID")

	parsedApplicationID, err := uuid.Parse(applicationID)
	if err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid ID"}
	}

	var application models.Application
	if err := initializers.DB.Preload("Opening").First(&application, "id = ?", parsedApplicationID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &fiber.Error{Code: 400, Message: "No Application of this ID found."}
		}
		return helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	if application.Opening.UserID.String() != loggedInUserID {
		return &fiber.Error{Code: 403, Message: "Do not have the permission to perform this action."}
	}

	if application.Status != 0 && application.Status != 1 {
		return &fiber.Error{Code: 400, Message: "Cannot perform this action now Now."}
	}

	if application.Status == 0 {
		application.Status = 1
	} else {
		application.Status = 0
	}
	result := initializers.DB.Save(&application)

	if result.Error != nil {
		return helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Application Under/Out of Review.",
	})
}