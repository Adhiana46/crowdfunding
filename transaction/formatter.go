package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt

	return formatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	var formatters []CampaignTransactionFormatter = []CampaignTransactionFormatter{}

	for _, transaction := range transactions {
		formatters = append(formatters, FormatCampaignTransaction(transaction))
	}

	return formatters
}

type UserTransactionFormatter struct {
	ID        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	userTransactionCampaignFormatter := CampaignFormatter{}
	userTransactionCampaignFormatter.Name = transaction.Campaign.Name

	if len(transaction.Campaign.CampaignImages) > 0 {
		for _, campaignImage := range transaction.Campaign.CampaignImages {
			if campaignImage.IsPrimary {
				userTransactionCampaignFormatter.ImageURL = campaignImage.FileName
			}
		}
	}

	formatter.Campaign = userTransactionCampaignFormatter

	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	var formatters []UserTransactionFormatter = []UserTransactionFormatter{}

	for _, transaction := range transactions {
		formatters = append(formatters, FormatUserTransaction(transaction))
	}

	return formatters
}
