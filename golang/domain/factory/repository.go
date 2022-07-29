package factory

import "github.com/rafaelbertuzzi/imersao-fullcycle-5/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
