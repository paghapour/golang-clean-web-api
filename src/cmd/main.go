package main

import (
	"github.com/paghapour/golang-clean-web-api/api"
	"github.com/paghapour/golang-clean-web-api/config"
	"github.com/paghapour/golang-clean-web-api/data/cache"
	"github.com/paghapour/golang-clean-web-api/data/db"
	"github.com/paghapour/golang-clean-web-api/pkg/logging"
)

func main() {
	cfg := config.GetConfig()
	
	logger := logging.NewLogger(cfg)
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil{
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = database.InitDb(cfg)
	defer database.CloseDb()
	if err != nil{
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}

	api.InitServer(cfg)
	
}
