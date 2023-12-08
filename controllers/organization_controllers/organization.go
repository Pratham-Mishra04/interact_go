package organization_controllers

import (
	"github.com/Pratham-Mishra04/interact/config"
	"github.com/Pratham-Mishra04/interact/helpers"
	"github.com/Pratham-Mishra04/interact/initializers"
	"github.com/Pratham-Mishra04/interact/models"
	API "github.com/Pratham-Mishra04/interact/utils/APIFeatures"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetOrganization(c *fiber.Ctx) error {
	orgID := c.Params("orgID")

	var organization models.Organization
	if err := initializers.DB.First(organization, "id=?", orgID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &fiber.Error{Code: 400, Message: "No Organization of this ID Found."}
		}
		return helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status":       "success",
		"message":      "",
		"organization": organization,
	})
}

func GetOrganizationTasks(c *fiber.Ctx) error {
	orgID := c.Params("orgID")

	var organization models.Organization
	if err := initializers.DB.
		Preload("Memberships").
		Preload("Memberships.User").
		Find(&organization, "id = ? ", orgID).Error; err != nil {
		return helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, Err: err}
	}

	var tasks []models.Task
	if err := initializers.DB.
		Preload("Users").
		Preload("SubTasks").
		Preload("SubTasks.Users").
		Find(&tasks, "organization_id = ? ", orgID).Error; err != nil {
		return helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status":       "success",
		"message":      "",
		"tasks":        tasks,
		"organization": organization,
	})
}

func GetOrganizationChats(c *fiber.Ctx) error {
	orgID := c.Params("orgID")

	var organization models.Organization
	if err := initializers.DB.
		Preload("Memberships").
		Preload("Memberships.User").
		Find(&organization, "id = ? ", orgID).Error; err != nil {
		return helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, Err: err}
	}

	var chats []models.GroupChat
	if err := initializers.DB.
		Preload("User").
		Preload("Memberships").
		Preload("Memberships.User").
		Find(&chats, "organization_id = ? ", orgID).Error; err != nil {
		return helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status":       "success",
		"message":      "",
		"chats":        chats,
		"organization": organization,
	})
}

func GetOrgEvents(c *fiber.Ctx) error {
	orgID := c.Params("orgID")

	paginatedDB := API.Paginator(c)(initializers.DB)

	var events []models.Event //TODO add last edited n all fields
	if err := paginatedDB.
		Preload("Organization").
		Where("organization_id = ?", orgID).
		Order("created_at DESC").
		Find(&events).Error; err != nil {
		return helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "",
		"events":  events,
	})
}

func GetOrganizationHistory(c *fiber.Ctx) error {
	orgID := c.Params("orgID")

	paginatedDB := API.Paginator(c)(initializers.DB)
	var history []models.OrganizationHistory

	if err := paginatedDB.
		Preload("User").
		Preload("Post").
		Preload("Event").
		Preload("Project").
		Preload("Task").
		Preload("Invitation").
		Where("organization_id=?", orgID).
		Order("created_at DESC").
		Find(&history).Error; err != nil {
		return helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, Err: err}
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"history": history,
	})
}