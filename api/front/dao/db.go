package dao

import (
	"flag"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	dns    string
	dirver string
	db     *sqlx.DB
)

func InitDB() {
	flag.StringVar(&dns, "db_dns", "liulang:root@tcp(127.0.0.1:3306)/bibirt", "database url")
	flag.StringVar(&dirver, "db_driver", "mysql", "database driver")
	db = sqlx.MustOpen(dirver, dns)
}
