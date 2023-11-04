package rdbadapter

import (
	"database/sql"
	"log/slog"
	"fmt"
	"net/url"
	_ "github.com/go-sql-driver/mysql"
)


type MySQLConfig struct {
    USERNAME string
	PASSWORD string
	HOST string
	PORT string
	NAME string
}

type MySQL struct {
	Conn *sql.DB
}

func NewMySQL(cfg *MySQLConfig) (*MySQL, error) {
    var err error
	conn, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=%s",
			cfg.USERNAME, cfg.PASSWORD, cfg.HOST, cfg.PORT, cfg.NAME, url.PathEscape("Asia/Tokyo")),
	)
	if err != nil {
		slog.Info("Fail to connect db", "error_detail", err.Error())
		return &MySQL{}, nil
	}
	slog.Info("Create DB connection.")
	return &MySQL{Conn: conn}, nil
}
