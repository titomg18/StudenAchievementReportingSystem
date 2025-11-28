package services

import (
	"errors"
	"time"

	"StudenAchievementReportingSystem/app/models"
	"StudenAchievementReportingSystem/app/repository"

	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
	Login(request *models.LoginRequest) (*models.LoginResponse, error)
	Register(request *models.RegisterRequest) error
}

type authService struct {
	userRepo *repository.UserRepository
	jwtKey   []byte
}

func NewAuthService(userRepo *repository.UserRepository, jwtSecret string) AuthService {
	return &authService{
		userRepo: userRepo,
		jwtKey:   []byte(jwtSecret),
	}
}

func (s *authService) Login(request *models.LoginRequest) (*models.LoginResponse, error) {

	// Pakai salah satu: email / username
	identifier := request.Email
	if identifier == "" {
		identifier = request.Username
	}

	// Cari user
	user, err := s.userRepo.GetUserByUsernameOrEmail(identifier)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	if !user.IsActive {
		return nil, errors.New("account is deactivated")
	}

	// Verifikasi password
	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password)) != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT
	token, err := s.generateToken(user)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	// Hapus password
	user.PasswordHash = ""

	// Response final
	return &models.LoginResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (s *authService) generateToken(user *models.User) (string, error) {
	expiration := time.Now().Add(24 * time.Hour)

	claims := &jwt.RegisteredClaims{
		Subject:   user.ID,
		ExpiresAt: jwt.NewNumericDate(expiration),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(s.jwtKey)
}

func (s *authService) Register(req *models.RegisterRequest) error {

	// hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	newUser := &models.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashed),
		FullName:     req.FullName,
		RoleID:       req.RoleID,
		IsActive:     true,
	}

	return s.userRepo.CreateUser(newUser)
}
