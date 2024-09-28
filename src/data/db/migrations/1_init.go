package migrations

import (
	"github.com/paghapour/golang-clean-web-api/config"
	db "github.com/paghapour/golang-clean-web-api/data/db"
	"github.com/paghapour/golang-clean-web-api/data/models"
	"github.com/paghapour/golang-clean-web-api/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

func Up1() {
	database := db.GetDb()

	tables := []interface{}{}

	country := models.Country{}
	city := models.City{}

	if !database.Migrator().HasTable(country) {
		tables = append(tables, country)
	}

	if !database.Migrator().HasTable(city) {
		tables = append(tables, city)
	}

	database.Migrator().CreateTable(tables...)
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)
}

func Down1() {}
