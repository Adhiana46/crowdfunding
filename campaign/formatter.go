package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	Slug             string `json:"slug"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

type CampaignDetailFormatter struct {
	ID               int                      `json:"id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	ImageURL         string                   `json:"image_url"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	BackerCount      int                      `json:"backer_count"`
	UserID           int                      `json:"user_id"`
	Description      string                   `json:"description"`
	User             CampaignUserFormatter    `json:"user"`
	Perks            []string                 `json:"perks"`
	Images           []CampaignImageFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

type CampaignImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{
		ID:               campaign.ID,
		Slug:             campaign.Slug,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageURL:         "",
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
	}

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignFormatters := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatters = append(campaignFormatters, FormatCampaign(campaign))
	}

	return campaignFormatters
}

func FormatCampaignDetail(c Campaign) CampaignDetailFormatter {
	f := CampaignDetailFormatter{
		ID:               c.ID,
		Name:             c.Name,
		ShortDescription: c.ShortDescription,
		ImageURL:         "",
		GoalAmount:       c.GoalAmount,
		CurrentAmount:    c.CurrentAmount,
		BackerCount:      c.BackerCount,
		UserID:           c.UserID,
		Description:      c.Description,
		User:             CampaignUserFormatter{},
		Perks:            []string{},
		Images:           []CampaignImageFormatter{},
	}

	// User
	if c.User.ID != 0 {
		f.User = CampaignUserFormatter{
			Name:      c.User.Name,
			AvatarURL: c.User.AvatarFileName,
		}
	}

	// Campaign Image
	if len(c.CampaignImages) > 0 {
		for _, campaignImage := range c.CampaignImages {
			if campaignImage.IsPrimary {
				f.ImageURL = campaignImage.FileName
			}

			imageFormatter := CampaignImageFormatter{
				ImageURL:  campaignImage.FileName,
				IsPrimary: campaignImage.IsPrimary,
			}
			f.Images = append(f.Images, imageFormatter)
		}
	}

	var perks []string = []string{}
	for _, perk := range strings.Split(c.Perks, ",") {
		if strings.TrimSpace(perk) != "" {
			perks = append(perks, strings.TrimSpace(perk))
		}
	}

	f.Perks = perks

	return f
}
