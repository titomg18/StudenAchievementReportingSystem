package services

import (
	"errors"
	"time"

	"StudenAchievementReportingSystem/app/models"
	"StudenAchievementReportingSystem/app/repository"

	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4"
)

// AuthService menangani logika bisnis untuk autentikasi
type AuthService interface {
	Login(request *models.LoginRequest) (*models.LoginResponse, error)
}

type authService struct {
	userRepo repository.UserRepository
	jwtKey   []byte
}

// NewAuthService membuat instance baru dari AuthService
func NewAuthService(userRepo repository.UserRepository, jwtSecret string) AuthService {
	return &authService{
		userRepo: userRepo,
		jwtKey:   []byte(jwtSecret),
	}
}

func (s *authService) Login(request *models.LoginRequest) (*models.LoginResponse, error) {
	// Cari user berdasarkan email atau username
	user, err := s.userRepo.GetUserByUsernameOrEmail(request.Email, request.Username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Periksa status aktif user
	if !user.IsActive {
		return nil, errors.New("account is deactivated")
	}

	// Verifikasi password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := s.generateToken(user)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	// Sembunyikan PasswordHash dalam response
	user.PasswordHash = ""

	response := &models.LoginResponse{
		Token: token,
		User:  []models.User{*user},
	}

	return response, nil
}

// generateToken membuat JWT token untuk user yang berhasil login
func (s *authService) generateToken(user *models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &jwt.RegisteredClaims{
		Subject:   user.ID,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtKey)
}