package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/lorenasg1/codepix/domain/model"
)

type PixKeyRepositoryDB struct {
	DB *gorm.DB
}

func (repository PixKeyRepositoryDB) AddBank(bank *model.Bank) error {

	err := repository.DB.Create(bank).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository PixKeyRepositoryDB) AddAccount(account *model.Account) error {

	err := repository.DB.Create(account).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository PixKeyRepositoryDB) RegisterKey(pixKey *model.PixKey) (*model.PixKey, error) {

	err := repository.DB.Create(pixKey).Error

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}

func (repository PixKeyRepositoryDB) FindKeyByKind(key string, kind string) (*model.PixKey, error) {

	var pixKey model.PixKey

	repository.DB.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("No key was found")
	}

	return &pixKey, nil
}

func (repository PixKeyRepositoryDB) FindAccount(id string) (*model.Account, error) {

	var account model.Account

	repository.DB.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, fmt.Errorf("No account was found")
	}

	return &account, nil
}

func (repository PixKeyRepositoryDB) FindBank(id string) (*model.Bank, error) {

	var bank model.Bank

	repository.DB.Preload("Bank").First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, fmt.Errorf("No bank was found")
	}

	return &bank, nil
}
