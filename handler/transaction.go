package handler

import (
	"bwastartup-api/helper"
	"bwastartup-api/transaction"
	"bwastartup-api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

// GET /api/v1/campaings/:id/transactions
func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput

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

	formatter := transaction.FormatCampaignTransactions(transactions)
	response := helper.APIResponse("Campaign transactions retrieved successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// GET /api/v1/user/:id/transactions
func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	var input transaction.GetUserTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get user transactions.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactions, err := h.service.GetTransactionsByUserID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get user transactions.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := transaction.FormatUserTransactions(transactions)
	response := helper.APIResponse("User transactions retrieved successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
