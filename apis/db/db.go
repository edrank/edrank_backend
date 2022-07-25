package db

import (
	"database/sql"
	"fmt"

	"github.com/edrank/edrank_backend/apis/types"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init(dbConfig *types.DatabaseConfig) {
	connString := fmt.Sprintf("%s:%s@(%s)/%s?parseTime=true", dbConfig.DbUser, dbConfig.DbPassword, dbConfig.DbHost, dbConfig.DbName)

	dbc, err := sql.Open("mysql", connString)
	DB = dbc

	if err != nil {
		panic(err)
	}
}

func GetDatabase() *sql.DB {
	return DB
}
