package main

import (
	"database/sql"
	"log/slog"
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
	logger *slog.Logger
}

func NewDB(cfg *Config, logger *slog.Logger) SqlDb {
    var err error
	conn, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=%s",
			cfg.USERNAME, cfg.PASSWORD, cfg.HOST, cfg.PORT, cfg.NAME, url.PathEscape("Asia/Tokyo")),
	)
	logger.Info("Create DB connection.")
	if err != nil {
		logger.Info("Fail to connect db", "error_detail", err.Error())
	}
    return SqlDb{conn, logger}
}
