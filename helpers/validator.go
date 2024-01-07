package helpers

import (
	"fmt"
	"strings"

	"github.com/Pratham-Mishra04/interact/schemas"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Validate[T any](payload T) error {

	validate := validator.New()

	if err := validate.Struct(payload); err != nil {

		validationErrors := err.(validator.ValidationErrors)

		var errorsBuilder strings.Builder

		for _, fieldError := range validationErrors {
			field := fieldError.Field()
			errorMessage := fmt.Sprintf("Invalid %s \n", field)
			errorsBuilder.WriteString(errorMessage)
		}

		errorsString := errorsBuilder.String()

		return &fiber.Error{Code: 400, Message: errorsString}
	}
	return nil
}

func ValidateReview(reqBody *schemas.ReviewReqBody) error {
    if len(reqBody.ReviewContent) <= 5 || len(reqBody.ReviewContent) >= 250 {
        return fiber.NewError(fiber.StatusBadRequest, "ReviewContent must be between 5 and 250 characters")
    }
    if reqBody.ReviewRating <= 1 || reqBody.ReviewRating >= 5 {
        return fiber.NewError(fiber.StatusBadRequest, "ReviewRating must be between 1 and 5")
    }
    return nil
}