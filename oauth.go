package fortnox

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v9"
	"golang.org/x/oauth2"
)

const (
	rdsTokenKey = "fortnox_token"
)

var (
	ErrNoTokenInTokenStorage = errors.New("fortnox: no token in token storage")
)

type TokenStorage interface {
	GetToken(ctx context.Context) (string, error)
	SetToken(ctx context.Context, token []byte) error
}

type tokenStorage struct {
	rc *redis.Client
}

func NewTokenStorage(rc *redis.Client) TokenStorage {
	return &tokenStorage{rc: rc}
}

func (t *tokenStorage) GetToken(ctx context.Context) (string, error) {
	tokenJSON, err := t.rc.Get(ctx, rdsTokenKey).Result()
	if err != nil && err == redis.Nil {
		return "", ErrNoTokenInTokenStorage
	}
	return tokenJSON, nil
}

func (t *tokenStorage) SetToken(ctx context.Context, token []byte) error {
	return t.rc.Set(ctx, rdsTokenKey, token, 0).Err()
}

type Oauth2Config struct {
	oauth2.Config
}

func NewOauth2Config() *Oauth2Config {
	config := &Oauth2Config{
		Config: oauth2.Config{
			RedirectURL:  "",
			ClientID:     "",
			ClientSecret: "",
			Scopes:       []string{},
			Endpoint: oauth2.Endpoint{
				AuthURL:   "https://apps.fortnox.se/oauth-v1/auth",
				TokenURL:  "https://apps.fortnox.se/oauth-v1/token",
				AuthStyle: oauth2.AuthStyleInHeader,
			},
		},
	}

	return config
}
