package handler

import (
	"bwastartup-api/campaign"
	"bwastartup-api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

// /api/v1/campaigns
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	var input campaign.GetCampaignsInput

	err := c.ShouldBindQuery(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get list campaigns.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaigns, err := h.service.GetCampaigns(input.UserID)
	if err != nil {
		response := helper.APIResponse("Error to get Campaigns.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatCampaigns(campaigns)
	response := helper.APIResponse("Campaigns retrieved successfully.", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// /api/v1/campaign/:id
func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail campaign.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// id, _ := strconv.Atoi(c.Param("id"))

	campaignObj, err := h.service.GetCampaign(input.ID)
	if err != nil {
		response := helper.APIResponse("Error to get Campaign.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatCampaignDetail(campaignObj)
	response := helper.APIResponse("Campaigns retrieved successfully.", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
