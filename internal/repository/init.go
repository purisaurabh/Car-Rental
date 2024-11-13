package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/purisaurabh/car-rental/internal/config"
)

func Init(ctx context.Context) (RepositoryStruct, error) {
	dbConfig := config.Database()
	db, err := sql.Open(dbConfig.Driver(), dbConfig.ConnectionURL())
	if err != nil {
		fmt.Println("Error in opening connection : ", err)
		return RepositoryStruct{}, err
	}

	if err = db.Ping(); err != nil {
		fmt.Println("Error in pinging the connection : ", err)
		return RepositoryStruct{}, err
	}

	db.SetMaxIdleConns(dbConfig.MaxPoolSize())
	db.SetMaxOpenConns(dbConfig.MaxOpenConns())
	db.SetConnMaxLifetime(time.Duration(dbConfig.MaxLifeTimeMins()) * time.Minute)
	return RepositoryStruct{DB: db}, nil
}
