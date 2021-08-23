package storage

import (
	"github.com/challenge/pkg/models"
	log "github.com/sirupsen/logrus"
)

func Migrate() {
	db := GetInstance()

	err := db.AutoMigrate(models.User{})

	if err != nil {
		log.Fatalf("error executing migrations %s", err.Error())
	}
}
