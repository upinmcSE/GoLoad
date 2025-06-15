package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/upinmcSE/GoLoad/internal/configs"

	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB(databaseConfig configs.Database) (*sql.DB, func(), error) {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		databaseConfig.Username,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Database,
	)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Printf("error connecting to the database: %+v\n", err)
		return nil, nil, err
	}
	
	cleanup := func() {
		db.Close()
	}

	return db, cleanup, nil
}


func InitGoquDB(db *sql.DB) (*goqu.Database, error) {
	return goqu.New("mysql", db), nil
}