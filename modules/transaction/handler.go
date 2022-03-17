package transaction

import (
	"bwastartup-api/helper"
	"bwastartup-api/modules/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service}
}

// GET /api/v1/campaings/:id/transactions
func (h *handler) GetCampaignTransactions(c *gin.Context) {
	var input GetCampaignTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign transactions.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// get from JWT
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionsByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign transactions.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := FormatCampaignTransactions(transactions)
	response := helper.APIResponse("Campaign transactions retrieved successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// GET /api/v1/transactions
func (h *handler) GetUserTransactions(c *gin.Context) {
	// get from JWT
	currentUser := c.MustGet("currentUser").(user.User)

	transactions, err := h.service.GetTransactionsByUserID(currentUser.ID)
	if err != nil {
		response := helper.APIResponse("Failed to get user transactions.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := FormatUserTransactions(transactions)
	response := helper.APIResponse("User transactions retrieved successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// POST /api/v1/transactions
func (h *handler) CreateTransaction(c *gin.Context) {
	var input CreateTransactionInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create transaction.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	input.User = c.MustGet("currentUser").(user.User)

	newTransaction, err := h.service.CreateTransaction(input)
	if err != nil {
		response := helper.APIResponse("Failed to create transaction.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := FormatTransaction(newTransaction)
	response := helper.APIResponse("Transaction created successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *handler) GetNotification(c *gin.Context) {
	var input TransactionNotificationInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Failed to process notification.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.service.ProcessPayment(input)
	if err != nil {
		response := helper.APIResponse("Failed to process notification.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, input)
}
