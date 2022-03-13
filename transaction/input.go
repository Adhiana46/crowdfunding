package transaction

import "bwastartup-api/user"

type GetUserTransactionsInput struct {
	ID int `uri:"id" binding:"required"`
}

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}