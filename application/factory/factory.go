package factory

import (
	"github.com/jinzhu/gorm"
	"github.com/lorenasg1/codepix/application/usecase"
	"github.com/lorenasg1/codepix/infra/repository"
)

func TransactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDB{DB: database}
	transactionRepository := repository.TransactionRepositoryDB{DB: database}

	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixRepository:         pixRepository,
	}

	return transactionUseCase
}
