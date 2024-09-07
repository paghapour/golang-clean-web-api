package main

import (
	"github.com/paghapour/golang-clean-web-api/api"
	"github.com/paghapour/golang-clean-web-api/config"
	"github.com/paghapour/golang-clean-web-api/data/cache"
)

func main() {
	cfg := config.GetConfig()

	cache.InitRedis(cfg)
	defer cache.CloseRedis()
	api.InitServer(cfg)
	
}
