package schemas

import "github.com/lib/pq"

type ProjectCreateSchema struct {
	Title       string         `json:"title" validate:"required,max=20"`
	Tagline     string         `json:"tagline" validate:"required,max=50"`
	Description string         `json:"description" validate:"required,max=1000"`
	Tags        pq.StringArray `json:"tags"`
	Category    string         `json:"category" validate:"required"`
	CoverPic    string         `json:"coverPic"`
	IsPrivate   bool           `json:"isPrivate" validate:"boolean"`
	Links       pq.StringArray `json:"links" validate:"dive,url"`
}

type ProjectUpdateSchema struct {
	Tagline      string         `json:"tagline" validate:"max=50"`
	Category     string         `json:"category"`
	CoverPic     string         `json:"coverPic" validate:"image"`
	Description  string         `json:"description" validate:"max=1000"`
	Page         string         `json:"page"`
	Tags         pq.StringArray `json:"tags"`
	IsPrivate    bool           `json:"isPrivate" validate:"boolean"`
	Links        pq.StringArray `json:"links" validate:"dive,url"`
	PrivateLinks pq.StringArray `json:"privateLinks" validate:"dive,url"`
}
