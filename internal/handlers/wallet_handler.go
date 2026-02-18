package handlers

import (
	// "github.com/Vladmir-dev/fintech-wallet/internal/models"
	"github.com/Vladmir-dev/fintech-wallet/internal/services"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

type WalletHandler struct {
	Service *services.WalletService
}

func NewWalletHandler(service *services.WalletService) *WalletHandler {
	return &WalletHandler{Service: service}
}

type DepositRequest struct {
	WalletId uint `json:"wallet_id" binding:"required"`
	Amount float64 `json:"amount" binding:"required,gt=0"`
	Reference string `json:"reference" binding:"required"`
}

type WithdrawRequest struct {
	WalletId uint `json:"wallet_id" binding:"required"`
	Amount float64 `json:"amount" binding:"required,gt=0"`
	Reference string `json:"reference" binding:"required"`
}

type TransferRequest struct {
	ToWalletID uint `json:"to_wallet_id" binding:"required"`
	FromWalletID uint `json:"from_wallet_id" binding:"required"`
	Amount float64 `json:"amount" binding:"required,gt=0"`
	Reference string `json:"reference" binding:"required"`
}

func (h *WalletHandler) Deposit(c *gin.Context) {
	//assuming wallet_id is passed as a URL parameter
	

	// Parse request body
	var body DepositRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	walletID := body.WalletId
	// walletID, err := strconv.ParseUint(walletIDStr, 10, 64)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid wallet ID"})
	// 	return
	// }

	// Call the service to perform the deposit
	if err := h.Service.Deposit(uint(walletID), body.Amount, body.Reference); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deposit successful"})
}


func (h *WalletHandler) Withdraw(c *gin.Context) {
	
	var body WithdrawRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	walletID := body.WalletId

	if err := h.Service.Withdraw(uint(walletID), body.Amount, body.Reference); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "withdraw successful"})
}


func (h *WalletHandler) Transfer(c *gin.Context) {
	// toWalletIDStr := c.Param("to_wallet_id")

	// toWalletID, err := strconv.ParseUint(toWalletIDStr, 10, 32)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid destination wallet ID"})
	// 	return
	// }

	var body TransferRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	toWalletID := body.ToWalletID
	fromWalletID := body.FromWalletID


	if err := h.Service.Transfer(fromWalletID, uint(toWalletID), body.Amount, body.Reference); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "transfer successful"})

}


//get transactions for a wallet id
func (h *WalletHandler) GetTransactions(c *gin.Context) {
   walletIDStr := c.Param("wallet_id")

	walletID, err := strconv.ParseUint(walletIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid wallet ID"})
		return
	}

	transactions, err := h.Service.GetTransactions(uint(walletID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}