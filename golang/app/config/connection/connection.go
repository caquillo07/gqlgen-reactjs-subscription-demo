package connection

import (
	"database/sql"

	"github.com/caquillo07/gqlgen-reactjs-subscription-demo/golang/app/config/constants"
	_ "github.com/go-sql-driver/mysql" // to get the db to compile
)

func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := constants.DBUSER
	dbPass := constants.DBPASSWORD
	dbHost := constants.DBHOST
	dbPort := constants.DBPORT
	dbName := constants.DBNAME
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}
