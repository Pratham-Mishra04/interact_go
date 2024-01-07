package routines

import (
	"math/rand"

	"github.com/Pratham-Mishra04/interact/helpers"
	"github.com/Pratham-Mishra04/interact/initializers"
	"github.com/Pratham-Mishra04/interact/models"
	"github.com/google/uuid"
)

func ComputeRelevance(reviewID uuid.UUID) {
	var review models.OrganizationReview
	if err := initializers.DB.First(&review, "id=?", reviewID).Error; err != nil {
		helpers.LogDatabaseError("No Post of this ID found-IncrementPostShare.", err, "go_routine")
	}
	review.Relevance = rand.Intn(91) + 10
	if err := initializers.DB.Save(&review).Error; err != nil {
		helpers.LogDatabaseError("Error in saving post-IncrementPostShare.", err, "go_routine")
	}
}