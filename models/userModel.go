package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	ID                        uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name                      string            `gorm:"varchar(25);not null" json:"name"`
	Username                  string            `gorm:"varchar(10);unique;not null" json:"username"`
	Email                     string            `gorm:"unique;not null" json:"email"`
	Password                  string            `json:"-"`
	ProfilePic                string            `gorm:"default:default.jpg" json:"profilePic"`
	CoverPic                  string            `gorm:"default:default.jpg" json:"coverPic"`
	PhoneNo                   string            `json:"phoneNo"`
	Bio                       string            `json:"bio"`
	Title                     string            `json:"title"`
	Tagline                   string            `json:"tagline"`
	Tags                      pq.StringArray    `gorm:"type:text[]" json:"tags"`
	PasswordResetToken        string            `json:"-"`
	PasswordResetTokenExpires time.Time         `json:"-"`
	Views                     int               `json:"views"` //! Show No of Views
	NoFollowing               int               `gorm:"default:0" json:"noFollowing"`
	NoFollowers               int               `gorm:"default:0" json:"noFollowers"`
	PasswordChangedAt         time.Time         `gorm:"default:current_timestamp" json:"-"`
	Admin                     bool              `gorm:"default:false" json:"-"`
	Active                    bool              `gorm:"default:true" json:"-"` //! add a functionality that on delete the acc goes inActive and if the user logs in within 30 days, it goes active again
	CreatedAt                 time.Time         `gorm:"default:current_timestamp" json:"-"`
	Achievements              []Achievement     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"achievements,omitempty"`
	Applications              []Application     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"applications,omitempty"`
	PostBookmarks             []PostBookmark    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"postBookmarks,omitempty"`
	ProjectBookmarks          []ProjectBookmark `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"projectBookmarks,omitempty"`
	Notifications             []Notification    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"notifications,omitempty"`
	SendNotifications         []Notification    `gorm:"foreignKey:SenderID;constraint:OnDelete:CASCADE" json:"sendNotifications,omitempty"`
} //! add last viewed projects

type ProfileView struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	UserID uuid.UUID `gorm:"type:uuid;not null" json:"userID"`
	Date   time.Time `gorm:"type:date" json:"date"`
	Count  int       `json:"count"`
}

// This func is called before gorm conversion
// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	u.PasswordChangedAt = time.Now()
// 	return nil
// }

type FollowFollower struct { //* follower follows followed
	FollowerID uuid.UUID `json:"followerID"`
	Follower   User      `gorm:"foreignKey:FollowerID" json:"follower"`
	FollowedID uuid.UUID `json:"followedID"`
	Followed   User      `gorm:"foreignKey:FollowedID" json:"followed"`
	CreatedAt  time.Time `gorm:"default:current_timestamp" json:"-"`
}

type Achievement struct {
	ID     uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	UserID uuid.UUID      `gorm:"type:uuid;not null" json:"userID"`
	User   User           `json:"user"`
	Title  string         `gorm:"type:text;not null" json:"title"`
	Skills pq.StringArray `gorm:"type:text[];not null" json:"skills"`
}
