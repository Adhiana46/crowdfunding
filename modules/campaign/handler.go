package campaign

import (
	"bwastartup-api/helper"
	"bwastartup-api/modules/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service}
}

// GET /api/v1/campaigns
func (h *handler) GetCampaigns(c *gin.Context) {
	var input GetCampaignsInput

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

	formatter := FormatCampaigns(campaigns)
	response := helper.APIResponse("Campaigns retrieved successfully.", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// GET /api/v1/campaign/:id
func (h *handler) GetCampaign(c *gin.Context) {
	var input GetCampaignDetailInput

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

	formatter := FormatCampaignDetail(campaignObj)
	response := helper.APIResponse("Campaigns retrieved successfully.", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// POST /api/v1/campaigns
func (h *handler) CreateCampaign(c *gin.Context) {
	var input CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create Campaign Failed.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// get from JWT
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.APIResponse("Create Campaign Failed.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := FormatCampaign(newCampaign)

	response := helper.APIResponse("Campaign has been created.", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

// PUT /api/v1/campaigns/:id
func (h *handler) UpdateCampaign(c *gin.Context) {
	var inputID GetCampaignDetailInput
	var inputData CreateCampaignInput
	var err error

	err = c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Update Campaign Failed.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Update Campaign Failed.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// get from JWT
	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	updatedCampaign, err := h.service.UpdateCampaign(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Update Campaign Failed.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := FormatCampaign(updatedCampaign)

	response := helper.APIResponse("Campaign has been updated.", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *handler) UploadImage(c *gin.Context) {
	var input CreateCampaignImageInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// get from JWT
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	userID := currentUser.ID
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image.", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.SaveCampaignImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image.", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Campaign image uploaded successfully.", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}
