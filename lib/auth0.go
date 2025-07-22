package lib

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

type IAuth0 interface {
	GetTokenClaims(string) (*Auth0TokenClaims, error)
}

type Auth0 struct {
	env Env
}

type Auth0UserInfo struct {
	Email      string `json:"email"`
	FamilyName string `json:"family_name"`
	GivenName  string `json:"given_name"`
	Identities []struct {
		Provider   string `json:"provider"`
		UserId     string `json:"user_id"`
		Connection string `json:"connection"`
		IsSocial   bool   `json:"isSocial"`
	} `json:"identities"`
	Name   string `json:"name"`
	UserId string `json:"user_id"`
}

type Auth0TokenClaims struct {
	RegisteredClaims struct {
		Iss string   `json:"iss"`
		Sub string   `json:"sub"`
		Aud []string `json:"aud"`
		Exp int64    `json:"exp"`
		Iat int64    `json:"iat"`
	} `json:"RegisteredClaims"`
	CustomClaims CustomClaims `json:"CustomClaims"`
}

type CustomClaims struct {
	Auth0UserInfo
}

type CombinedClaims struct {
	validator.ValidatedClaims
	CustomClaims
}

func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

func (a *Auth0) GetTokenClaims(token string) (*Auth0TokenClaims, error) {
	issuerURL, err := url.Parse(a.env.AUTH0_DOMAIN)
	if err != nil {
		return nil, err
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{a.env.AUTH0_CLIENT_ID},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		return nil, err
	}

	jsonClaims, err := jwtValidator.ValidateToken(context.TODO(), token)
	if err != nil {
		splittedErrMsg := strings.Split(err.Error(), ":")

		return nil, errors.New(strings.TrimSpace(splittedErrMsg[len(splittedErrMsg)-1]))
	}

	validatedClaims, ok := jsonClaims.(*validator.ValidatedClaims)
	if !ok {
		return nil, errors.New("type assertion failed")
	}

	// Now you can convert the validatedClaims to []byte (if that's what you need)
	jsonBytes, err := json.Marshal(validatedClaims)
	if err != nil {
		return nil, err
	}

	var auth0TokenClaims Auth0TokenClaims
	err = json.Unmarshal([]byte(jsonBytes), &auth0TokenClaims)
	if err != nil {
		return nil, err
	}

	return &auth0TokenClaims, nil
}

func NewAuth0(env Env) IAuth0 {
	return &Auth0{env: env}
}
