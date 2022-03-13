package transaction

import (
	"bwastartup-api/campaign"
	"bwastartup-api/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Campaign   campaign.Campaign
	User       user.User
}
