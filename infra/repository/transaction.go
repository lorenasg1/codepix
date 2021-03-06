package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/lorenasg1/codepix/domain/model"
)

type TransactionRepositoryDB struct {
	DB *gorm.DB
}

func (t *TransactionRepositoryDB) Register(transaction *model.Transaction) error {
	err := t.DB.Create(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionRepositoryDB) Save(transaction *model.Transaction) error {
	err := t.DB.Save(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionRepositoryDB) Find(id string) (*model.Transaction, error) {

	var transaction model.Transaction

	t.DB.Preload("AccountFrom.Bank").First(&transaction, "id: ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("No transaction was found")
	}

	return &transaction, nil
}
