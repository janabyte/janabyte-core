package service

import (
	"fmt"
	"github.com/aidosgal/janabyte/janabyte-core/internal/http/model"
	"github.com/aidosgal/janabyte/janabyte-core/internal/http/token"
	"net/http"
	"strings"
)

func (s *UserService) AuthenticateUser(r *http.Request, user *model.User) (*model.User, error) {
	const op = "service.AuthenticateUserByLogin"
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		fields := strings.Fields(authHeader)
		if len(fields) != 2 || fields[0] != "Bearer" {
			return nil, fmt.Errorf("invalid authorization header")
		}
		tk := fields[1]
		claims, err := token.VerifyToken(tk)
		if err != nil {
			return nil, fmt.Errorf("invalid token")
		}
		_, err = s.GetUserById(claims.Id)
		if err != nil {
			return nil, fmt.Errorf("invalid user")
		}
		return nil, nil
	}
	if user.Email != "" && user.Login == "" {
		err := s.repository.AuthenticateByEmail(user.Email, user.Password)
		if err != nil {
			return nil, fmt.Errorf("error authenticating with email %s: %s", user.Email, err)
		}
		retUser, err := s.repository.GetUserByEmail(user.Email)
		if err != nil {
			return nil, fmt.Errorf("error authenticating with email %s: %s", user.Email, err)
		}
		return retUser, nil
	}
	if user.Login != "" && user.Email == "" {
		err := s.repository.AuthenticateByLogin(user.Login, user.Password)
		if err != nil {
			return nil, fmt.Errorf("error authenticating with login %s: %s", user.Login, err)
		}
		retUser, err := s.repository.GetUserByLogin(user.Login)
		if err != nil {
			return nil, fmt.Errorf("error authenticating with login %s: %s", user.Login, err)
		}
		return retUser, nil
	}
	return nil, fmt.Errorf("invalid user")
}
