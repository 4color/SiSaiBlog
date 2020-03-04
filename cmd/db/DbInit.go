package db

import (
	"database/sql"
	"github.com/4color/SiSaiBlog/configs"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var SqlDB *sql.DB

func init() {
	var err error
	//SqlDB, err = sql.Open(strings.ToLower(goracle.DefaultConnectionClass), config.Config.App.Db.ConnectionString)
	SqlDB, err = sql.Open("sqlite3", configs.Config.App.Db.ConnectionString)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
