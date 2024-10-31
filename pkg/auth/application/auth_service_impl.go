package application

import (
    "errors"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/auth/domain/valueobjects"
    "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/domain/entities"
    "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/domain/repositories"
    "golang.org/x/crypto/bcrypt"
)

type authService struct {
    userRepo repositories.UserRepository
    jwtKey   []byte
}

func NewAuthService(userRepo repositories.UserRepository, jwtKey []byte) AuthService {
    return &authService{
        userRepo: userRepo,
        jwtKey:   jwtKey,
    }
}

func (s *authService) Register(user entities.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    return s.userRepo.Create(user)
}

func (s *authService) Login(credentials valueobjects.Credentials) (string, error) {
    user, err := s.userRepo.GetByEmail(credentials.Email)
    if err != nil {
        return "", errors.New("user not found")
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
    if err != nil {
        return "", errors.New("invalid credentials")
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, err := token.SignedString(s.jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}