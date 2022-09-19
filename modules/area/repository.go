package area

import (
	"database/sql"

	"gorm.io/gorm"
)

type Repository interface {
	Save(area Area) (Area, error)
}
type repository struct {
	db       *gorm.DB
	dbManual *sql.DB
}

func NewRepository(db *gorm.DB, dbManual *sql.DB) *repository {
	return &repository{db, dbManual}
}

func (r *repository) Save(area Area) (Area, error) {
	err := r.db.Create(&area).Error
	if err != nil {
		return area, err
	}

	return area, nil
}
