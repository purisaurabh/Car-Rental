package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/go-sql-driver/mysql"
	"github.com/purisaurabh/car-rental/internal/pkg/constants"
	"go.uber.org/zap"
)

type RepositoryStruct struct {
	DB *sql.DB
}

type Repository interface {
	UserRegistration(ctx context.Context, userReq *UserRegistrationRepo) (int64, error)
	UserLogin(ctx context.Context, email string) (UserInfo, error)
}

func NewRepository(DB *sql.DB) Repository {
	return &RepositoryStruct{
		DB: DB,
	}
}

func (user *RepositoryStruct) UserRegistration(ctx context.Context, userReq *UserRegistrationRepo) (int64, error) {
	values := []interface{}{userReq.Name, userReq.Email, userReq.Password, userReq.Mobile, userReq.Role, userReq.CreatedAt, userReq.UpdatedAt}

	// using squirrel to insert data into the database
	insertQuery, args, err := squirrel.Insert("users").Columns(constants.UserRegistrationColumns...).Values(values...).PlaceholderFormat(squirrel.Question).ToSql()
	if err != nil {
		zap.S().Error("error generating the user insert query : ", err)
		return 0, err
	}

	result, err := user.DB.ExecContext(ctx, insertQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.S().Error("error inserting the user data : ", err)
			return 0, err
		}

		if err, ok := err.(*mysql.MySQLError); ok {
			if err.Number == 1062 {
				zap.S().Error("error inserting the user data : ", err)
				return 0, err
			}
		}

		zap.S().Error("error inserting the user data : ", err)
		return 0, err
	}
	userID, err := result.LastInsertId()
	if err != nil {
		zap.S().Error("error getting the last inserted user id : ", err)
		return 0, err
	}
	return userID, nil
}

func (user *RepositoryStruct) UserLogin(ctx context.Context, email string) (UserInfo, error) {
	var userInfo UserInfo
	query, args, err := squirrel.Select(constants.UserLoginColumns...).From("users").Where(squirrel.Eq{"email": email}).ToSql()
	if err != nil {
		zap.S().Error("error generating the user login query : ", err)
		return UserInfo{}, err
	}

	err = user.DB.QueryRowContext(ctx, query, args...).Scan(&userInfo.UserID, &userInfo.Name, &userInfo.Password, &userInfo.Mobile, &userInfo.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.S().Error("error getting the user data : ", err)
			return UserInfo{}, fmt.Errorf("user not found")
		}
		zap.S().Error("error getting the user data : ", err)
		return UserInfo{}, err
	}

	return userInfo, nil

}
