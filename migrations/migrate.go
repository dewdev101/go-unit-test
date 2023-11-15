package migrations

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB, model interface{}) {
	if err := db.AutoMigrate(model); err != nil {
		panic("unable to migrate model")
	}
}
