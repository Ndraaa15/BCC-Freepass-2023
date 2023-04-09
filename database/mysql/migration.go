package database

import (
	"bcc-freepass-2023/src/entity"

	"gorm.io/gorm"
)

type dbMigrate struct {
	db *gorm.DB
}

func (m *dbMigrate) AutoMigrate() error {
	if err := m.db.AutoMigrate(
		&entity.Student{},
	); err != nil {
		return err
	}

	return nil
}
