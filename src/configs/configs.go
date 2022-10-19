package configs

import (
	"os"

	"github.com/chmenegatti/k8s-study/src/utils"
	"github.com/joho/godotenv"
)

func GetEnvKey(key string) (envKey string, err error) {
	err = godotenv.Load(".env")

	if err != nil {
		utils.Logger("Erro", "GetEnvKey", "Erro ao abrir .env")
		return "", err
	}

	envKey = os.Getenv(key)

	return envKey, err
}
