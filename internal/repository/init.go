package repository

import (
	"context"
	"time"

	"github.com/purisaurabh/car-rental/internal/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(ctx context.Context) (db *gorm.DB, err error) {
	dbConfig := config.Database()
	dsn := dbConfig.ConnectionURL()

	db, err = gorm.Open(mysql.Open(dsn) , &gorm.Config{})
	if err != nil {
		zap.S().Errorw("Error in opening the connection", "error", err)
		return nil, err
	}

	sqlDB , err := db.DB()	
	if err != nil {
		zap.S().Errorw("Error in getting the database connection", "error", err)
		return nil, err
	}

	// ping the database
	if err = sqlDB.PingContext(ctx); err != nil {
		zap.S().Errorw("Error in pinging the database", "error", err)
		return nil, err
	}


	sqlDB.SetMaxIdleConns(dbConfig.MaxPoolSize())
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns())
	sqlDB.SetConnMaxLifetime(time.Duration(dbConfig.MaxLifeTimeMins()) * time.Minute)
	return db, nil
}
