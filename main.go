package main

import (
	"self-payrol/config"
	"sync"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Infof(".env is not loaded properly")
	}
	log.Infof("read .env from file")

	config := config.NewConfig()
	defer closeDB(config.Database())
	server := InitServer(config)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Run()
	}()

	wg.Wait()
}

func closeDB(DB *gorm.DB) {
	sqldb, err := DB.DB()
	if err != nil {
		log.Infof("failed to close db %s", err)
	}

	err = sqldb.Close()
	if err != nil {
		log.Infof("failed to close db %s", err)
	}
}
