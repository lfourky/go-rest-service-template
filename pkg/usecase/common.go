package usecase

import (
	"github.com/lfourky/go-rest-service-template/pkg/model/domain"
	"github.com/lfourky/go-rest-service-template/pkg/repository"
	"github.com/lfourky/go-rest-service-template/pkg/service/log"
)

func rollback(tx repository.Store, logger *log.Logger) {
	if err := tx.Rollback(); err != nil {
		logger.WithError(err).Error("unable to rollback a transaction")
	}
}

func beginTx(store repository.Store, logger *log.Logger) (repository.Store, error) {
	tx, err := store.BeginTransaction()
	if err != nil {
		logger.WithError(err).Warn("unable to start a transaction")

		return nil, ErrDatabaseInternal
	}

	return tx, nil
}

func commit(tx repository.Store, logger *log.Logger) error {
	if err := tx.Commit(); err != nil {
		logger.WithError(err).Warn("unable to commit a transaction")

		return ErrDatabaseInternal
	}

	return nil
}

func findUserByID(id string, tx repository.Store, logger *log.Logger) (*domain.User, error) {
	userID, err := domain.ParseUUID(id)
	if err != nil {
		return nil, ErrDomainInvalidUUID
	}

	user, err := tx.User().FindByID(userID)
	if err != nil {
		logger.WithError(err).Warn("unable to find user")

		return nil, ErrDatabaseInternal
	}

	if user == nil {
		return nil, ErrDatabaseUserNotFound
	}

	return user, nil
}
