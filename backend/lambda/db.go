package main

import (
	"database/sql"
	"fmt"
	"net/url"
	_ "github.com/go-sql-driver/mysql"
)


type Config struct {
    USERNAME string
	PASSWORD string
	HOST string
	PORT string
	NAME string
}

type SqlDb struct {
	conn *sql.DB
}

func NewDB(cfg *Config) SqlDb {
    var err error
	conn, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=%s",
			cfg.USERNAME, cfg.PASSWORD, cfg.HOST, cfg.PORT, cfg.NAME, url.PathEscape("Asia/Tokyo")),
	)
	if err != nil {
		fmt.Println("Fail to connect db" + err.Error())
	}
    return SqlDb{conn}
}
