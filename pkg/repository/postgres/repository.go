package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/lfourky/go-rest-service-template/pkg/repository"
	"github.com/lfourky/go-rest-service-template/pkg/service/log"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Repository struct {
	db *gorm.DB

	user *User
	item *Item
}

func New(cfg Config, logger *log.Logger) (repository.Store, error) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.DatabaseName,
	)), &gorm.Config{
		Logger: gormLogger(cfg, logger),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	return newWithDB(db), nil
}

func newWithDB(db *gorm.DB) *Repository {
	return &Repository{
		db: db,

		user: &User{db},
		item: &Item{db},
	}
}

func (r *Repository) BeginTransaction() (repository.Store, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return newWithDB(tx), nil
}

func (r *Repository) Commit() error {
	return r.db.Commit().Error
}

func (r *Repository) Rollback() error {
	if err := r.db.Rollback().Error; !errors.Is(err, sql.ErrTxDone) {
		return err
	}

	return nil
}

func (r *Repository) User() repository.User {
	return r.user
}

func (r *Repository) Item() repository.Item {
	return r.item
}

func gormLogger(cfg Config, logger *log.Logger) glogger.Interface {
	level := glogger.Silent

	if cfg.LogMode {
		switch logger.Level {
		case logrus.TraceLevel, logrus.DebugLevel, logrus.InfoLevel:
			level = glogger.Info
		case logrus.WarnLevel:
			level = glogger.Warn
		case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
			level = glogger.Error
		}
	}

	return glogger.New(
		logger,
		glogger.Config{
			SlowThreshold: time.Second,
			LogLevel:      level,
			Colorful:      false,
		},
	)
}
