package app

import "github.com/revel/revel"
import "github.com/jinzhu/gorm"
import _ "github.com/go-sql-driver/mysql"
import _ "github.com/mattn/go-sqlite3"

var DB *gorm.DB

func InitDB() *gorm.DB {
    var (
        driver string
        spec   string
        found  bool
    )

    // Read the configuration.
    if driver, found = revel.Config.String("db.driver"); !found {
        revel.ERROR.Fatal("Database driver configuration not set! (db.driver)")
    }

    if spec, found = revel.Config.String("db.spec"); !found {
        revel.ERROR.Fatal("Database spec configuration not set! (db.spec)")
    }

    maxIdleConns := revel.Config.IntDefault("db.max_idle_conns", 10)
    maxOpenConns := revel.Config.IntDefault("db.max_open_conns", 100)
    singularTable := revel.Config.BoolDefault("db.singular_table", false)
    logMode := revel.Config.BoolDefault("db.log_mode", false)

    // Initialise gorm
    dbm, err := gorm.Open(driver, spec)
    if err != nil {
        revel.ERROR.Fatal(err)
    }

    DB = &dbm

    dbm.DB().Ping()
    dbm.DB().SetMaxIdleConns(maxIdleConns)
    dbm.DB().SetMaxOpenConns(maxOpenConns)
    dbm.SingularTable(singularTable)
    dbm.LogMode(logMode)
    dbm.SetLogger(gorm.Logger{revel.INFO})

    return &dbm
}
