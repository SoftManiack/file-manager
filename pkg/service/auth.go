package service

import (
	"crypto/sha1"
	"errors"
	user "file-manager"
	"file-manager/pkg/repository"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

const (
	salt       = "hjgrhjqw126417ajfha"
	signingKey = "qrkjk#4#%352ol57&2dl"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

type tokenClaims struct {
	jwt.StandardClaims
	UserUid string `json:"user_uid"`
}

func (s *AuthService) CreateUser(user user.User) (string, error) {
	user.Password = generatePasswordHash(user.Password)

	fmt.Println(user.Name)
	uid, err := s.repo.CreateUser(user)

	if err == nil {
		var path = fmt.Sprintf("C:/Програмирование/Проекты/file-manager/cloud/%s", uid)

		err = os.Mkdir(path, 0775)

		if err != nil {
			return "", err
		}
	}

	return uid, err
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetUser(email, generatePasswordHash(password))
	logrus.Errorln(err)
	if err != nil {
		return " ", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Uid,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserUid, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
