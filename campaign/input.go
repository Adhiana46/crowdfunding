package campaign

type GetCampaignsInput struct {
	UserID int `form:"user_id"`
}

type GetCampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
