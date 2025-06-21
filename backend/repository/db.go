package repository

import (
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDB() (*sqlx.DB, error) {
	config := mysql.Config{
		Net:                  "tcp",
		User:                 os.Getenv("NS_MARIADB_USER"),
		Passwd:               os.Getenv("NS_MARIADB_PASSWORD"),
		Addr:                 os.Getenv("NS_MARIADB_HOST") + ":" + os.Getenv("NS_MARIADB_PORT"),
		DBName:               os.Getenv("NS_MARIADB_DATABASE"),
		ParseTime:            true,
		AllowNativePasswords: true,
		Collation:            "utf8mb4_general_ci",
	}
	return sqlx.Connect("mysql", config.FormatDSN())
}
