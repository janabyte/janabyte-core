package token

import (
	"fmt"
	"github.com/aidosgal/janabyte/janabyte-core/internal/http/model"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

const SecretKey = "Secret Key"

func MakeToken(user *model.User, permissions []string, duration time.Duration) (string, *Claims, error) {
	const op = "token.MakeToken"
	claims, err := CreationUserClaims(user, permissions, duration)
	if err != nil {
		return "", nil, fmt.Errorf("%s: %s", op, err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", nil, fmt.Errorf("error creating token %s: %s", op, err)
	}
	return tokenStr, claims, nil
}

func VerifyToken(tokenStr string) (*Claims, error) {
	const op = "token.VerifyToken"
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v %s", token.Header["alg"], op)

		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token %s: %s", op, err)
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("error parsing token %s: %s", op, err)
	}
	return claims, nil
}

func AccessTokenCookie(r *http.Request) (*Claims, error) {
	const op = "accessTokenCookie"
	cookie, err := r.Cookie("access_token")
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	claims, err := VerifyToken(cookie.Value)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	return claims, nil
}
func RefreshTokenCookie(r *http.Request) (*Claims, error) {
	const op = "token.CheckAuthByCookie"
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		return nil, fmt.Errorf("%s - %s", op, err)
	}
	claims, err := VerifyToken(cookie.Value)
	if err != nil {
		return nil, fmt.Errorf("%s - %s", op, err)
	}
	return claims, nil
}
