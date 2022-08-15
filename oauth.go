package fortnox

import (
	"golang.org/x/oauth2"
)

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
