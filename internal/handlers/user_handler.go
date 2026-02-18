package handlers

import (
"github.com/Vladmir-dev/fintech-wallet/internal/services"
"github.com/Vladmir-dev/fintech-wallet/internal/models"
"net/http"
"github.com/gin-gonic/gin"
)

type OnboardRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	Currency  string `json:"currency" binding:"required"`
}

type UserHandler struct {
   Service * services.UserService	// UserService *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) OnboardUser(c *gin.Context)  {
	var body OnboardRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	
	user := models.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  body.Password,
	}
	created, err := h.Service.CreateUser(user, body.Currency)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"user": created,
	})

	// return h.Service.CreateUser(user, "USD")
}