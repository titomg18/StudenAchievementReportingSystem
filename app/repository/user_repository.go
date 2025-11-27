package repository

import (
	"database/sql"
	"StudenAchievementReportingSystem/app/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// func (r *UserRepository) GetUserByUsernameOrEmail(identifier string) (*models.User, error) {
// 	// Implementation sementara
// 	return &models.User{
// 		ID:           "test-id",
// 		Username:     "testuser",
// 		Email:        "test@example.com",
// 		PasswordHash: "$2a$10$examplehash", // Ganti dengan hash bcrypt yang valid
// 		FullName:     "Test User",
// 		RoleID:       "role-id",
// 		IsActive:     true,
// 	}, nil
// }

func (r *UserRepository) GetUserByUsernameOrEmail(email, username string) (*models.User, error) {
	return &models.User{
		ID:           "test-id",
		Username:     "testuser",
		Email:        "test@example.com",
		PasswordHash: "$2a$10$examplehash", // Ganti dengan hash bcrypt yang valid
		FullName:     "Test User",
		RoleID:       "role-id",
		IsActive:     true,
	}, nil
}

func (r *UserRepository) GetUserPermissions(roleID string) ([]string, error) {
	// Implementation sementara
	return []string{"read", "write"}, nil
}