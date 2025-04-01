package configs

import "github.com/mxpadidar/letsgo/internal/infra/services"

type Configs struct {
	PostgresDSN    string
	JWTSecret      []byte
	AccessTokenTTL int // time to live in seconds, for example, 3600
	BcryptCost     int // cost factor for bcrypt hashing, for example, 10
}

func Load(envService *services.EnvVarService) (*Configs, error) {
	if err := envService.LoadEnvironmentFile(); err != nil {
		return nil, err
	}

	configs := &Configs{}
	var err error

	if configs.PostgresDSN, err = envService.GetString("POSTGRES_DSN"); err != nil {
		return nil, err
	}
	if configs.JWTSecret, err = envService.GetBytes("JWT_SECRET"); err != nil {
		return nil, err
	}
	if configs.AccessTokenTTL, err = envService.GetInt("ACCESS_TOKEN_TTL"); err != nil {
		return nil, err
	}
	if configs.BcryptCost, err = envService.GetInt("BCRYPT_COST"); err != nil {
		return nil, err
	}
	return configs, nil
}
