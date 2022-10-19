package migrations

import (
	"github.com/chmenegatti/k8s-study/src/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Nomes{})
}
