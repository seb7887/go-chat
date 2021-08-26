package storage

import (
	"sync"

	log "github.com/challenge/pkg/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func GetInstance() *gorm.DB {
	// Singleton pattern to have only one global DB instance
	once.Do(func() {
		database, err := gorm.Open(sqlite.Open("chat.db"), &gorm.Config{})
		if err != nil {
			log.Fatalf("error initializing SQLite3 database %s", err.Error())
		}
		DB = database
	})

	return DB
}
