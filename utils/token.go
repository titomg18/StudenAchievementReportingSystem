package utils

import (
    "errors"
    "os"
    "time"
    "StudenAchievementReportingSystem/app/models/postgresql"
    "StudenAchievementReportingSystem/config"
    "github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *models.User, roleName string, permissions []string) (string, error) {
    jwtCfg := config.LoadJWT()

    claims := &models.JWTClaims{
        UserID:      user.ID,
        RoleID:      user.RoleID,
        RoleName:    roleName,
        Permissions: permissions,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtCfg.TTLHours) * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            Issuer:    "student-achievement-system",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtCfg.Secret)
}

func ValidateToken(tokenString string) (*models.JWTClaims, error) {
    jwtCfg := config.LoadJWT()
    
    token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
        return jwtCfg.Secret, nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(*models.JWTClaims); ok && token.Valid {
        return claims, nil
    }
    
    return nil, errors.New("invalid token claims")
}

func GenerateRefreshToken(user *models.User) (string, error) {
    secret := os.Getenv("JWT_REFRESH_SECRET")
    if secret == "" {
        secret = os.Getenv("JWT_SECRET")
        if secret == "" {
            secret = "default_refresh_secret"
        }
    }

    expiration := time.Now().Add(7 * 24 * time.Hour)

    claims := &models.RefreshClaims{
        UserID: user.ID.String(),
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expiration),
            Issuer:    "student-achievement-system",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}

func ValidateRefreshToken(tokenString string) (*models.RefreshClaims, error) {
    secret := os.Getenv("JWT_REFRESH_SECRET")
    if secret == "" {
        secret = os.Getenv("JWT_SECRET")
        if secret == "" {
            secret = "default_refresh_secret"
        }
    }

    token, err := jwt.ParseWithClaims(tokenString, &models.RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })

    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*models.RefreshClaims)
    if !ok || !token.Valid {
        return nil, errors.New("invalid refresh token")
    }

    return claims, nil
}