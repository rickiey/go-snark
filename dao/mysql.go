package dao

import (
	"database/sql"
	"go-snark/conf"

	"fmt"
	"net/url"

	// ..
	_ "github.com/go-sql-driver/mysql"
)

var (
	// DB ..
	DB *sql.DB
)

// InitDB ..
func InitDB() {
	var err error

	dbconf := conf.Conf.Db
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=%s&parseTime=true", dbconf.Uname, dbconf.Passwd, dbconf.Server,
		dbconf.Port, dbconf.Dbname, url.QueryEscape("Asia/Shanghai"))
	DB, err = sql.Open("mysql", dns)

	if nil != err {
		panic(err)
	}
	DB.SetMaxOpenConns(dbconf.MaxConns)

	err = DB.Ping()
	if nil != err {
		panic(err)
	}
}
