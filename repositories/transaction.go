package repositories

import (
	"testdumpflix/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	CreateTransaction(category models.Transaction) (models.Transaction, error)
	UpdateTransaction(status string, ID string) error
	DeleteTransaction(category models.Transaction) (models.Transaction, error)
	GetOneTransaction(ID string) (models.Transaction, error)
}

type repositoryTransaction struct {
	db *gorm.DB
}

func RepositoryforTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("User").Find(&transactions).Error

	return transactions, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Debug().Preload("User").First(&transaction, ID).Error

	return transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Debug().Create(&transaction).Error

	return transaction, err
}

func (r *repository) UpdateTransaction(status string, ID string) error {
	var transaction models.Transaction
	r.db.Preload("User").First(&transaction, ID)

	// If is different & Status is "success" decrement product quantity
	if status != transaction.Status && status == "success" {
		var user models.User
		r.db.First(&user, transaction.User.ID)
		user.Subscribe = true
		r.db.Save(&user)
	}

	transaction.Status = status

	err := r.db.Save(&transaction).Error

	return err
}

func (r *repository) GetOneTransaction(ID string) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("User.Profile").First(&transaction, "id = ?", ID).Error

	return transaction, err
}

func (r *repository) DeleteTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error // Using Delete method

	return transaction, err
}
