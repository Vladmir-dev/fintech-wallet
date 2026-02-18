package services

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/Vladmir-dev/fintech-wallet/internal/models"
	// "github.com/go-playground/locales/currency"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) CreateUser(req models.User, currency string) (*models.User, error) {

	var existingUser models.User
	if err := s.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return nil, gorm.ErrRegistered
	}

	// Hash the password
	 hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	 req.Password = string(hashed)

	 //create wallet in a transaction
	 err := s.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&req).Error; err != nil {
			return err
		}

		wallet := models.Wallet{
			UserID:   req.ID,
			Currency: currency,
			Balance:  0,
		}

		if err := tx.Create(&wallet).Error; err != nil {
			return err
		}

		// req.Wallet = wallet
		return nil
	 })

	 if err != nil {
		return nil, err
	 }

	return &req, nil
}

func (s *UserService) UserProfile(userID uint) (*models.User, error) {
	var user models.User

	//check that user exists
	if err := s.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}

	//load wallet details
	
	if err := s.DB.Preload("Wallet").First(&user, userID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}