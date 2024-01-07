package schemas

type ReviewReqBody struct {
	ReviewContent string `json:"review_content"`
	ReviewRating  int    `json:"review_rating"`
	Anonymous     bool   `json:"anonymous"`
}
