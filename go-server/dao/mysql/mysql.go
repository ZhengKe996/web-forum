package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go-web-example/settings"
)

var db *sqlx.DB

func InitDB(config *settings.MYSQLConfig) (err error) {
	// DSN:Data Source Name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", config.User, config.Password, config.Host, config.Port, config.DbName)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	return
}

func Close() {
	db.Close()
}
