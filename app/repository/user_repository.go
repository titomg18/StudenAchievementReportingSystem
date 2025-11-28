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

// Hanya pakai 1 identifier: email ATAU username
func (r *UserRepository) GetUserByUsernameOrEmail(identifier string) (*models.User, error) {
	query := `
		SELECT 
			id,
			username,
			email,
			password_hash,
			full_name,
			role_id,
			is_active,
			created_at,
			updated_at
		FROM users
		WHERE email = $1 OR username = $1
		LIMIT 1;
	`

	row := r.db.QueryRow(query, identifier)

	var user models.User

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.FullName,
		&user.RoleID,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Tidak ditemukan
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserPermissions(roleID string) ([]string, error) {
	// nanti bisa ambil dari DB; sementara return dummy
	return []string{"read", "write"}, nil
}

// func NewUserRepository(db *sql.DB) *UserRepository {
// 	return &UserRepository{db: db}
// }

// func (r *UserRepository) GetUserByUsernameOrEmail(email, username string) (*models.User, error) {
// 	query := `
// 		SELECT 
// 			id,
// 			username,
// 			email,
// 			password_hash,
// 			full_name,
// 			role_id,
// 			is_active,
// 			created_at,
// 			updated_at
// 		FROM users
// 		WHERE email = $1 OR username = $2
// 		LIMIT 1;
// 	`

// 	row := r.db.QueryRow(query, email, username)

// 	var user models.User

// 	err := row.Scan(
// 		&user.ID,
// 		&user.Username,
// 		&user.Email,
// 		&user.PasswordHash,
// 		&user.FullName,
// 		&user.RoleID,
// 		&user.IsActive,
// 		&user.CreatedAt,
// 		&user.UpdatedAt,
// 	)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, nil // user tidak ditemukan
// 		}
// 		return nil, err
// 	}

// 	return &user, nil
// }

// func (r *UserRepository) GetUserPermissions(roleID string) ([]string, error) {
// 	return []string{"read", "write"}, nil
// }
