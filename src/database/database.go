package database

import (
	"fmt"
	"os"
	"time"

	"github.com/chmenegatti/k8s-study/src/configs"
	"github.com/chmenegatti/k8s-study/src/database/migrations"
	"github.com/chmenegatti/k8s-study/src/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDb() {
	info := utils.Trace()
	var (
		user, _ = configs.GetEnvKey("DB_USER")
		pass, _ = configs.GetEnvKey("DB_PASS")
		name, _ = configs.GetEnvKey("DB_NAME")
		url, _  = configs.GetEnvKey("DB_URL")
		port, _ = configs.GetEnvKey("DB_PORT")
	)

	urlPort := fmt.Sprintf("%s:,%s", url, port)

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", url, port, user, name, pass)
	utils.Logger("Info", "ConnectDb", "Conectando no banco em: "+urlPort)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		utils.Logger("Erro", info, "Erro ao conectar ao Banco de Dados")
		os.Exit(2)
	}

	utils.Logger("Info", "ConnectDb", "Conectado em: "+urlPort+" com sucesso")

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)

}

func GetDatabase() *gorm.DB {
	return db
}
