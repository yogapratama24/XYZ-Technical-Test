package repository

import (
	"log"
	"xyz_test/model"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction model.Transaction) error
	ReadTransaction(userId int) (transactions []model.Transaction, err error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) CreateTransaction(transaction model.Transaction) error {
	db := r.db
	err := db.Create(&transaction)
	if err.Error != nil {
		log.Printf("Error create data transaction with err: %v", err)
		return err.Error
	}
	return nil
}

func (r *transactionRepository) ReadTransaction(userId int) (transactions []model.Transaction, err error) {
	db := r.db
	if err := db.Where("user_id = ?", userId).Joins("User").Find(&transactions).Error; err != nil {
		log.Printf("Token not found with err: %s", err)
		return transactions, err
	}
	return transactions, nil
}
