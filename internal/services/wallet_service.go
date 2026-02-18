package services

import (
	"errors"
	"github.com/Vladmir-dev/fintech-wallet/internal/models"
	"gorm.io/gorm"
	"fmt"
)

type WalletService struct {
	db *gorm.DB
}

func NewWalletService(db *gorm.DB) *WalletService {
	return &WalletService{db: db}
}

// get wallet by wallet id
func (s *WalletService) GetWalletByID(userID uint) (*models.Wallet, error) {
	var wallet models.Wallet
	if err := s.db.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}

// Deposit funds into a wallet
func (s *WalletService) Deposit(walletID uint, amount float64, reference string) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	return s.db.Transaction(func(tx *gorm.DB) error {

		//check for existing transaction with this reference
		var existingTxn models.Transaction

		if err := tx.Where("reference = ?", reference).First(&existingTxn).Error; err == nil {
			if existingTxn.WalletID == walletID && existingTxn.Amount == amount && existingTxn.Type == "deposit" {
				return nil
			}

			return errors.New("transaction with this reference already exists")
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		//create new transaction record
		txnNew := models.Transaction{
			WalletID:  walletID,
			Amount:    amount,
			Type:      "deposit",
			Reference: reference,
			//  Counterparty: "external",
		}

		if err := tx.Create(&txnNew).Error; err != nil {
			return err
		}

		// Update wallet balance
		if err := tx.Model(&models.Wallet{}).Where("id = ?", walletID).Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return err
		}

		return nil

	})
}

// Withdraw funds from a wallet
func (s *WalletService) Withdraw(walletID uint, amount float64, reference string) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
    return s.db.Transaction(func(tx *gorm.DB) error {

		//check for existing transaction with this reference
		var existingTransaction models.Transaction
		if err :=  tx.Where("reference = ?" , reference).First(&existingTransaction).Error; err == nil {
			if existingTransaction.WalletID == walletID && existingTransaction.Amount == amount && existingTransaction.Type == "withdrawal" {
				return nil
			}

			return errors.New("transaction with this reference already exists")
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		//check balance
		var wallet models.Wallet
		if err := tx.Where("id = ?", walletID).First(&wallet).Error; err != nil {
			return err
		}

		if wallet.Balance < amount {
			return errors.New("insufficient funds")
		}

		//create new transaction record
		txnNew := models.Transaction{
			WalletID:  walletID,
			Amount:    amount,
			Type:      "withdrawal",
			Reference: reference,
			// Counterparty: "external",
		}
		
		if err := tx.Create(&txnNew).Error; err != nil {
			return err
		}

		return nil

	})
}


func (s *WalletService) Transfer(fromWalletID uint, toWalletID uint, amount float64, reference string) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	if fromWalletID == toWalletID {
		return errors.New("cannot transfer to the same wallet")
	}

	return s.db.Transaction(func(tx *gorm.DB) error { 
		//check for existing transaction with this reference
		var existingTransaction models.Transaction
		if err := tx.Where("reference = ?", reference).First(&existingTransaction).Error; err == nil {
			if existingTransaction.WalletID == fromWalletID && existingTransaction.Amount == amount && existingTransaction.Type == "transfer_out" {
				return nil
			} 
			return errors.New("transaction with this reference already exists")
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		//check balance of from wallet
		var fromWallet models.Wallet
		if err := tx.Where("id = ?", fromWalletID).First(&fromWallet).Error; err != nil {
			return err
		}

		if fromWallet.Balance < amount {
			return errors.New("insufficient funds")
		}

		//create transaction record for from wallet
		txnFrom := models.Transaction{
			WalletID:  fromWalletID,
			Amount:    amount,
			Type:      "transfer_out",
			Reference: reference,
			Counterparty: fmt.Sprintf("wallet_%d", toWalletID),
		}

		if err := tx.Create(&txnFrom).Error; err != nil {
			return err
		}
		
		//create transaction record for to wallet
		txnTo := models.Transaction{
			WalletID:  toWalletID,
			Amount:    amount,
			Type:      "transfer_in",
			Reference: reference,
			Counterparty: fmt.Sprintf("wallet_%d", fromWalletID),
		}

		if err := tx.Create(&txnTo).Error; err != nil {
			return err
		}

		// Update balances
		if err := tx.Model(&models.Wallet{}).Where("id = ?", fromWalletID).Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
			return err
		}

		if err := tx.Model(&models.Wallet{}).Where("id = ?", toWalletID).Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return err
		}

		return nil

	})

}
