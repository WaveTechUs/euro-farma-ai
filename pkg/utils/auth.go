package utils

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type CustomClaims struct {
	UserRole string
	jwt.RegisteredClaims
}

func Encryption(data string) ([]byte, error) {
	cData, err := bcrypt.GenerateFromPassword([]byte(data), 14)
	return cData, err
}

func JwtClaims(userId int, userRole string) *jwt.Token {
	customClaims := CustomClaims{
		userRole,
		jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(userId)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)

	return claims
}

func SetCookie(token string, w http.ResponseWriter) {

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		Secure:   true,
		Domain:   "localhost:5173",
		SameSite: http.SameSiteNoneMode,
	}

	http.SetCookie(w, &cookie)
}

func JwtAuth(r *http.Request, restricted bool) error {
	cookie := jwtauth.TokenFromCookie(r)

	claims := &CustomClaims{}
	token, err := jwt.ParseWithClaims(cookie, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil || !token.Valid {
		return err
	}
	err = restrictedEndPointVerification(restricted, claims.UserRole)
	return err
}

func restrictedEndPointVerification(restricted bool, role string) error {
	if !restricted {
		return nil
	}

	if role == "admin" {
		return nil
	}

	return errors.New("access denied: you don't have the required permissions")
}
