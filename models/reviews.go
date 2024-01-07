package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type OrganizationReview struct {
	ID                  uuid.UUID             `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	UserID              uuid.UUID             `gorm:"type:uuid;not null" json:"userID"`
	User                User                  `gorm:"" json:"user"`
	OrganizationID      uuid.UUID             `gorm:"type:uuid;not null" json:"organizationID"`
	Organization        Organization          `gorm:"" json:"organization"`
	Review              string                `gorm:"type:text;not null" json:"review"`
	Relevance 			int 				  `gorm:"not null;default:10" json:"relevance"`	
	Rating 			    int                   `gorm:"not null" json:"rating"`
	LikeCount 	     	int            		  `gorm:"not null;default:0" json:"likeCount"`
	LikedBy    			pq.StringArray 		  `gorm:"type:uuid[]" json:"-"`
	DislikeCount 		int            		  `gorm:"not null;default:0" json:"dislikeCount"`
	DislikedBy 			pq.StringArray 		  `gorm:"type:uuid[]" json:"-"`
	Anonymous 		 	bool                  `gorm:"not null;default:false" json:"anonymous"`
	CreatedAt   		time.Time  			  `gorm:"default:current_timestamp" json:"createdAt"`
}
