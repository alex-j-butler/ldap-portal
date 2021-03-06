package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"

	"qixalite.com/Ranndom/ldap-portal/modules/settings"
	"qixalite.com/Ranndom/ldap-portal/modules/logging"
)

var DB *gorm.DB

func NewContext() *gorm.DB {
	dbm, err := gorm.Open(settings.Database.Driver, settings.Database.Spec)
	if err != nil {
		logging.Logger.Fatalf("Database error: %s", err)
	}

	DB = dbm

	dbm.DB().Ping()
	dbm.DB().SetMaxIdleConns(settings.Database.MaxIdleConns)
	dbm.DB().SetMaxOpenConns(settings.Database.MaxOpenConns)
	dbm.LogMode(settings.Database.LogMode)

	return dbm
}

func RegisterModel(model interface{}) {
	DB.AutoMigrate(model)
}

