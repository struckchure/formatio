package lib

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type IJwt interface {
	GenerateJWT(sub string) (*AuthTokens, error)
	VerifyJWT(tokenString string, jwtType JwtType) (*jwt.Token, error)
}

type Jwt struct {
	env Env
}

type AuthTokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type JwtType int

const (
	ACCESS_TOKEN_TYPE JwtType = iota + 1
	REFRESH_TOKEN_TYPE
)

func (t JwtType) String() string {
	return [...]string{"ACCESS_TOKEN", "REFRESH_TOKEN"}[t-1]
}

func (t JwtType) EnumIndex() int {
	return int(t)
}

func (j *Jwt) GenerateJWT(sub string) (*AuthTokens, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 6).Unix(),
	})
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	signedAccessToken, _ := accessToken.SignedString([]byte(j.env.JWT_ACCESS_KEY))
	signedRefreshToken, _ := refreshToken.SignedString([]byte(j.env.JWT_REFRESH_KEY))

	return &AuthTokens{
		AccessToken:  signedAccessToken,
		RefreshToken: signedRefreshToken,
	}, nil
}

func (j *Jwt) VerifyJWT(tokenString string, jwtType JwtType) (*jwt.Token, error) {
	var jwtSecret string

	switch jwtType {
	case ACCESS_TOKEN_TYPE:
		jwtSecret = j.env.JWT_ACCESS_KEY
	case REFRESH_TOKEN_TYPE:
		jwtSecret = j.env.JWT_REFRESH_KEY
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func NewJwt(env Env) IJwt {
	return &Jwt{env: env}
}
