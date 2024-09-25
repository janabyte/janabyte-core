package token

import (
	"fmt"
	"github.com/aidosgal/janabyte/janabyte-core/internal/http/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

type Claims struct {
	Id          int      `json:"id"`
	Login       string   `json:"login"`
	Email       string   `json:"email"`
	RoleId      int      `json:"role_id"`
	Permissions []string `json:"permissions"`
	jwt.StandardClaims
}

func CreationUserClaims(user *model.User, permissions []string, duration time.Duration) (*Claims, error) {
	const op = "token.CreationUserClaims"
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("error creating tokenID: %v %s", err, op)
	}
	return &Claims{
		Id:          user.Id,
		Login:       user.Login,
		Email:       user.Email,
		RoleId:      user.RoleID,
		Permissions: permissions,
		StandardClaims: jwt.StandardClaims{
			Id:        tokenId.String(),
			Subject:   fmt.Sprintf("%d", user.Id),
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}, nil
}
