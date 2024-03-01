package models

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	ID                uuid.UUID                `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	UserID            uuid.UUID                `gorm:"type:uuid;not null" json:"userID"` //user model who is given the organization status
	User              User                     `gorm:"" json:"user"`
	OrganizationTitle string                   `gorm:"unique" json:"title"`
	Memberships       []OrganizationMembership `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE" json:"memberships"`
	Invitations       []Invitation             `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE" json:"invitations"`
	History           []OrganizationHistory    `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE" json:"history"`
	Events            []Event                  `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE" json:"events"`
	Reviews           []Review                 `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE" json:"-"`
	ResourceBucket    []ResourceBucket         `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE" json:"-"`
	NumberOfMembers   int16                    `gorm:"default:0" json:"noMembers"`
	NumberOfEvents    int16                    `gorm:"default:0" json:"noEvents"`
	NumberOfProjects  int16                    `gorm:"default:0" json:"noProjects"`
	CreatedAt         time.Time                `gorm:"default:current_timestamp" json:"createdAt"`
}

type OrganizationRole string

const (
	Member  OrganizationRole = "Member"
	Senior  OrganizationRole = "Senior"
	Manager OrganizationRole = "Manager"
)

//* Member can only view.
//* Manager can CRUD on Projects and Posts.
//* Owner can CRUD on members can change their roles and can update the Organization Details.
//* Organization Account Logger can do all this, including delete the organization, transfer of ownership, and organization chats.

type OrganizationMembership struct {
	ID             uuid.UUID        `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	OrganizationID uuid.UUID        `gorm:"type:uuid;not null" json:"organizationID"`
	Organization   Organization     `gorm:"" json:"organization"`
	UserID         uuid.UUID        `gorm:"type:uuid;not null" json:"userID"`
	User           User             `gorm:"" json:"user"`
	Title          string           `gorm:"type:varchar(25);not null" json:"title"`
	Role           OrganizationRole `gorm:"type:text" json:"role"`
	CreatedAt      time.Time        `gorm:"default:current_timestamp" json:"createdAt"`
}

//TODO add last edited n all fields on models

/*
history type:
*-1 - Organization created
*0 - User created an event
*1 - User deleted an event
*2 - User updated an event
*3 - User invited a member
*4 - User withdraw an invitation
*5 - User removed a member
*6 - User made a post
*7 - User deleted a post
*8 - User edited a post
*9 - User added a project
*10 - User deleted a project
*11 - User edited a project
*12 - User added a task
*13 - User deleted a task
*14 - User edited Org Details
*15 - Member left Org
*16 - Event coordinators added
*17 - Event coordinators removed
*18 - User added a poll
*19 - User deleted a poll
*20 - User edited a poll
*21 - User added an announcement
*22 - User deleted an announcement
*23 - User edited an announcement
*24 - User added an opening
*25 - User deleted an opening
*26 - User edited an opening
*27 - User accepted an application
*28 - User rejected application
*29 - User withdraw an invitation for event cohost
*/

type OrganizationHistory struct {
	ID             uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	OrganizationID uuid.UUID    `gorm:"type:uuid;not null" json:"orgID"`
	HistoryType    int8         `json:"historyType"`
	UserID         uuid.UUID    `gorm:"type:uuid;not null" json:"userID"`
	User           User         `json:"user"`
	PostID         *uuid.UUID   `gorm:"type:uuid" json:"postID"`
	Post           Post         `json:"post"`
	EventID        *uuid.UUID   `gorm:"type:uuid" json:"eventID"`
	Event          Event        `json:"event"`
	ProjectID      *uuid.UUID   `gorm:"type:uuid" json:"projectID"`
	Project        Project      `json:"project"`
	TaskID         *uuid.UUID   `gorm:"type:uuid" json:"taskID"`
	Task           Task         `json:"task"`
	InvitationID   *uuid.UUID   `gorm:"type:uuid" json:"invitationID"`
	Invitation     Invitation   `json:"invitation"`
	AnnouncementID *uuid.UUID   `gorm:"type:uuid" json:"announcementID"`
	Announcement   Announcement `json:"announcement"`
	OpeningID      *uuid.UUID   `gorm:"type:uuid" json:"openingID"`
	Opening        Opening      `json:"opening"`
	ApplicationID  *uuid.UUID   `gorm:"type:uuid" json:"applicationID"`
	Application    Application  `json:"application"`
	PollID         *uuid.UUID   `gorm:"type:uuid" json:"pollID"`
	Poll           Poll         `json:"poll"`
	DeletedText    string       `gorm:"type:text" json:"deletedText"`
	CreatedAt      time.Time    `gorm:"default:current_timestamp" json:"createdAt"`
}
