package middlewares

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"

	"github.com/li-go/jwt-auth-service/models"
)

var (
	userKey             = "user"
	ErrInvalidToken     = errors.New("invalid token")
	ErrUserDataNotFound = errors.New("user data not found")
)

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINGKEY")), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

func toJWTToken(u *models.User) *jwt.Token {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = u.ID
	claims["name"] = u.Name
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	return token
}

func fromJWTToken(token *jwt.Token) (*models.User, error) {
	claims := token.Claims.(jwt.MapClaims)
	id, ok := claims["sub"].(float64)
	if !ok {
		return nil, ErrInvalidToken
	}
	name, ok := token.Claims.(jwt.MapClaims)["name"].(string)
	if !ok {
		return nil, ErrInvalidToken
	}
	return &models.User{ID: int(id), Name: name}, nil
}

func AuthHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := jwtMiddleware.CheckJWT(w, r); err != nil {
			return
		}

		user, err := fromJWTToken(r.Context().Value(userKey).(*jwt.Token))
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userKey, user)))
	})
}

func ToSignedToken(user *models.User) (string, error) {
	return toJWTToken(user).SignedString([]byte(os.Getenv("SIGNINGKEY")))
}

func FromRequest(r *http.Request) (*models.User, error) {
	user, ok := r.Context().Value(userKey).(*models.User)
	if !ok {
		return nil, ErrUserDataNotFound
	}
	return user, nil
}
