package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Vladmir-dev/fintech-wallet/internal/handlers"
	"github.com/Vladmir-dev/fintech-wallet/internal/services"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	// User routes
	router.POST("/users/register", userHandler.OnboardUser)
	
	// Wallet routes (to be implemented

	// Wallet routes (add these once you create WalletHandler)
	// walletService := services.NewWalletService(db)
	// walletHandler := handlers.NewWalletHandler(walletService)

	// router.POST("/wallet/deposit/:wallet_id", walletHandler.Deposit)
	// router.POST("/wallet/withdraw/:wallet_id", walletHandler.Withdraw)
	// router.POST("/wallet/transfer/:wallet_id", walletHandler.Transfer)
	// router.GET("/wallet/transactions/:wallet_id", walletHandler.GetTransactions)
	// router.GET("/user/profile/:user_id", userHandler.GetProfile)
}