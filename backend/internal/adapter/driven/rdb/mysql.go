package rdbadapter

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-sql-driver/mysql"
)

type MySQLConfig struct {
	USERNAME string
	PASSWORD string
	HOST     string
	PORT     string
	NAME     string
}

type MySQL struct {
	Conn *sql.DB
}

func NewMySQL(cfg *MySQLConfig) (*MySQL, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		slog.Info("Fail to load timezone", "error_detail", err.Error())
		return nil, NewRdbUnexpectedError(err)
	}

	dsnConfig := mysql.Config{
		DBName:               cfg.NAME,
		User:                 cfg.USERNAME,
		Passwd:               cfg.PASSWORD,
		Addr:                 fmt.Sprintf("%s:%s", cfg.HOST, cfg.PORT),
		Net:                  "tcp",
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  jst,
		AllowNativePasswords: true,
	}
	conn, err := sql.Open("mysql", dsnConfig.FormatDSN())
	if err != nil {
		slog.Info("Fail to open db", "error_detail", err.Error())
		return nil, NewRdbUnexpectedError(err)
	}
	err = conn.Ping()
	if err != nil {
		slog.Info("Fail to connect db", "error_detail", err.Error())
		return nil, NewRdbUnexpectedError(err)
	}
	slog.Info("Create DB connection.")
	return &MySQL{Conn: conn}, nil
}
