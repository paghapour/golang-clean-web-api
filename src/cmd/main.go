package main

import (
	"github.com/paghapour/golang-clean-web-api/api"
	"github.com/paghapour/golang-clean-web-api/config"
	"github.com/paghapour/golang-clean-web-api/data/cache"
	"github.com/paghapour/golang-clean-web-api/data/db"
)

func main() {
	cfg := config.GetConfig()

	cache.InitRedis(cfg)

	defer cache.CloseRedis()
	db.InitDb(cfg)
	defer db.CloseDb()
	
	api.InitServer(cfg)
	
}
