package repository

import "database/sql"

type RepositoryStruct struct {
	DB *sql.DB
}

type Repository interface {
}
