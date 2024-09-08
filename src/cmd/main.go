package main

import (
	"log"

	"github.com/paghapour/golang-clean-web-api/api"
	"github.com/paghapour/golang-clean-web-api/config"
	"github.com/paghapour/golang-clean-web-api/data/cache"
	"github.com/paghapour/golang-clean-web-api/data/db"
)

func main() {
	cfg := config.GetConfig()
	
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil{
		log.Fatal(err)
	}

	err = database.InitDb(cfg)
	defer database.CloseDb()
	if err != nil{
		log.Fatal(err)
	}

	api.InitServer(cfg)
	
}
