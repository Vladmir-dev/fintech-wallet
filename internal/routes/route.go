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
	walletService := services.NewWalletService(db)
	walletHandler := handlers.NewWalletHandler(walletService)

	// User routes
	router.POST("/users/onboard", userHandler.OnboardUser)
	router.GET("/user/profile/:user_id", userHandler.GetProfile)
	
	router.POST("/wallet/deposit", walletHandler.Deposit)
	router.POST("/wallet/withdraw", walletHandler.Withdraw)
	router.POST("/wallet/transfer", walletHandler.Transfer)
	router.GET("/wallet/transactions/:wallet_id", walletHandler.GetTransactions)
	
}