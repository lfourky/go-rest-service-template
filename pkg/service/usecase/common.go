package usecase

import (
	"database/sql"
	"log"

	"github.com/lfourky/go-rest-service-template/pkg/repository"
)

func rollbackOnError(tx repository.Store, txErr error) {
	if txErr != nil {
		if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
			log.Printf("couldn't rollback a transaction: %s", err)
		}
	}
}
