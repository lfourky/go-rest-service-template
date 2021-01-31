package domain

import (
	"database/sql/driver"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PrimaryKey represents a primary key for an entity in the form of an UUID.
// The reason it's needed is because not all UUIDs need to have GORM's BeforeCreate hook method.
// PrimaryKey encapsulates that behaviour, instead.
type PrimaryKey struct {
	ID UUID `gorm:"column:id"`
}

func (p *PrimaryKey) BeforeCreate(tx *gorm.DB) error {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	p.ID = UUID(newUUID.String())

	return nil
}

type UUID string

func (u UUID) Value() (driver.Value, error) {
	if u == "" {
		return nil, nil
	}

	val, err := uuid.Parse(string(u))
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (u *UUID) Scan(value interface{}) error {
	val, err := uuid.Parse(value.(string))
	if err != nil {
		return err
	}

	*u = UUID(val.String())

	return nil
}

func (u UUID) String() string {
	return string(u)
}

func NewUUID() (UUID, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return UUID(uuid.String()), nil
}

func ParseUUID(s string) (UUID, error) {
	_, err := uuid.Parse(s)
	if err != nil {
		return "", err
	}

	return UUID(s), nil
}
