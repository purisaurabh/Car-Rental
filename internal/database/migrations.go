package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/purisaurabh/car-rental/internal/config"
)

func RunMigrations() error {
	dbConfig := config.Database()
	db, err := sql.Open(dbConfig.Driver(), dbConfig.ConnectionURL())
	if err != nil {
		fmt.Println("Error in opening connection : ", err)
		return err
	}

	driver, err := getDBDriverInstance(db, dbConfig.Driver())
	if err != nil {
		fmt.Println("Error in getting driver instance : ", err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(GetMigrationPath(), dbConfig.Driver(), driver)
	if err != nil {
		fmt.Println("Error in creating migration instance : ", err)
		return err
	}

	err = m.Up()
	if err == nil && err == migrate.ErrNoChange {
		fmt.Println("No migration changes")
		return nil
	}

	return err
}

func CreateMigrationFile(fileName string) error {
	if len(fileName) == 0 {
		return errors.New("please provide a file name")
	}

	// Ensure the migrations directory exists
	migrationPath := config.MigrationPath()

	if err := os.MkdirAll(migrationPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create migrations directory: %w", err)
	}

	timeStamp := time.Now().Unix()

	upMigrationFilePath := fmt.Sprintf("%s/%d_%s.up.sql", migrationPath, timeStamp, fileName)
	downMigrationFilePath := fmt.Sprintf("%s/%d_%s.down.sql", migrationPath, timeStamp, fileName)

	if err := createFile(upMigrationFilePath); err != nil {
		fmt.Println("Error in creating up migration file: ", err)
		return err
	}

	fmt.Printf("Up migration file created at %s\n", upMigrationFilePath)

	if err := createFile(downMigrationFilePath); err != nil {
		fmt.Println("Error in creating down migration file: ", err)
		return err
	}

	fmt.Printf("Down migration file created at %s\n", downMigrationFilePath)

	return nil
}

func RollbackMigration(s string) error {
	steps, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error in converting steps to integer:", err)
		return err
	}

	if steps < 0 {
		return fmt.Errorf("negative steps not allowed: %d", steps)
	}

	connectionURL := "mysql_user:mysql_user@tcp(localhost:3306)/testDB?parseTime=true&loc=Local"
	m, err := migrate.New(GetMigrationPath(), "mysql://"+connectionURL)
	if err != nil {
		fmt.Println("Error in creating migration instance:", err)
		return err
	}

	err = m.Steps(-1 * steps)
	if err == nil || err == migrate.ErrNoChange {
		fmt.Println("Rollback successful")
		return nil
	}

	return err
}

func getDBDriverInstance(db *sql.DB, driver string) (database.Driver, error) {
	switch driver {
	case "mysql":
		return mysql.WithInstance(db, &mysql.Config{})
	default:
		return nil, errors.New("no migrate driver instance found")
	}
}

func createFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error in creating file : ", err)
		return err
	}

	err = file.Close()
	if err != nil {
		fmt.Println("Error in closing file", err)
		return err
	}
	return nil
}

func GetMigrationPath() string {
	return fmt.Sprintf("file://%s", config.MigrationPath())
}
