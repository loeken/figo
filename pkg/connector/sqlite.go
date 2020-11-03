package connector

import (
	"os"

	"github.com/loeken/figo/pkg/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnectSqlite to the database
func ConnectSqlite() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("sqlite_database.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if os.Getenv("DB_MIGRATE") == "true" {
		Migrate(db)
	}

	return db
}

// MigrateSqlite tables
func MigrateSqlite(db *gorm.DB) {
	db.Migrator().AutoMigrate(&entity.Release{})
	db.Migrator().AutoMigrate(&entity.Content{})
}
