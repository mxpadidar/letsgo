package configs

import (
	"github.com/mxpadidar/letsgo/internal/infra/helpers"
)

type Configs struct {
	PostgresDSN    string
	JWTSecret      []byte
	AccessTokenTTL int // time to live in seconds, for example, 3600
	BcryptCost     int // cost factor for bcrypt hashing, for example, 10
}

func InitConfigs() *Configs {
	return &Configs{
		PostgresDSN:    helpers.MustEnv("POSTGRES_DSN"),
		JWTSecret:      []byte(helpers.MustEnv("JWT_SECRET")),
		AccessTokenTTL: helpers.MustIntEnv("ACCESS_TOKEN_TTL"),
		BcryptCost:     helpers.MustIntEnv("BCRYPT_COST"),
	}
}
