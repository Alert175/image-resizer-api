package config

import (
	"image-resizer-api/pkg/helpers/logger"
	"image-resizer-api/src/adapters/http"
	"image-resizer-api/src/db/psql"
	"log"

	"github.com/joho/godotenv"
)

func Init() error {
	defer http.Init()
	envInit()
	psql.DbInit()
	startBackgroundProcess()

	return nil
}

func envInit() {
	if err := godotenv.Load("./config/.env"); err != nil {
		log.Fatalf("Error loading .env file")
	} else {
		logger.Success("env loaded")
	}
}

// запуск фоновых процессов
func startBackgroundProcess() {
	// logger.Log("started background process")
}
