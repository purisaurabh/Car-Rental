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
	return &RepositoryStruct{DB: DB}
}

func (repo *RepositoryStruct) UserRegistration(ctx context.Context, userReq *UserRegistrationRepo) (int64, error) {
	// Build the SQL query using squirrel
	query, args, err := squirrel.Insert(constants.UsersTable).
		Columns(constants.UserRegistrationColumns...).
		Values(userReq.Name, userReq.Email, userReq.Password, userReq.Mobile, userReq.Role, userReq.CreatedAt, userReq.UpdatedAt).
		PlaceholderFormat(squirrel.Question).
		ToSql()
	if err != nil {
		zap.S().Errorw("failed to build insert query", "error", err)
		return 0, fmt.Errorf("failed to build insert query: %w", err)
	}

	// Execute the query
	result, err := repo.DB.ExecContext(ctx, query, args...)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			zap.S().Warnw("duplicate entry for user registration", "email", userReq.Email, "error", err)
			return 0, fmt.Errorf("duplicate entry: %w", err)
		}
		zap.S().Errorw("failed to execute user registration query", "error", err)
		return 0, fmt.Errorf("failed to execute query: %w", err)
	}

	// Get the last inserted ID
	userID, err := result.LastInsertId()
	if err != nil {
		zap.S().Errorw("failed to get last inserted ID", "error", err)
		return 0, fmt.Errorf("failed to fetch last insert id: %w", err)
	}

	return userID, nil
}

func (repo *RepositoryStruct) UserLogin(ctx context.Context, email string) (UserInfo, error) {
	var userInfo UserInfo

	// Build the SQL query using squirrel
	query, args, err := squirrel.Select(constants.UserLoginColumns...).
		From(constants.UsersTable).
		Where(squirrel.Eq{"email": email}).
		PlaceholderFormat(squirrel.Question).
		ToSql()
	if err != nil {
		zap.S().Errorw("failed to build login query", "email", email, "error", err)
		return UserInfo{}, fmt.Errorf("failed to build login query: %w", err)
	}

	// Execute the query and scan the result
	err = repo.DB.QueryRowContext(ctx, query, args...).Scan(&userInfo.UserID, &userInfo.Name, &userInfo.Password, &userInfo.Mobile, &userInfo.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.S().Infow("user not found during login", "email", email)
			return UserInfo{}, fmt.Errorf("user not found")
		}
		zap.S().Errorw("failed to execute login query", "email", email, "error", err)
		return UserInfo{}, fmt.Errorf("failed to execute login query: %w", err)
	}

	return userInfo, nil
}
