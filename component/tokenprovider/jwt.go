package tokenprovider

import (
	"food/common"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtProvider struct {
	secret string
}

type myClaims struct {
	Payload TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func NewJwt(secret string) *jwtProvider {
	return &jwtProvider{
		secret: secret,
	}
}

func (j *jwtProvider) Generate(data TokenPayload, expiry int) (*Token, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		data, jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	})
	token, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}
	return &Token{
		Token:   token,
		Expiry:  expiry,
		Created: time.Now(),
	}, nil
}

func (j *jwtProvider) Validate(token string) (*TokenPayload, error) {
	t, err := jwt.ParseWithClaims(token, &myClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, common.ErrJWT(err)
	}
	if !t.Valid {
		return nil, common.ErrJWT(err)
	}
	claims, ok := t.Claims.(*myClaims)
	if !ok {
		return nil, common.ErrJWT(err)
	}
	return &claims.Payload, nil
}
