package configs

import (
	"github.com/mxpadidar/letsgo/internal/infra/helpers"
)

type Configs struct {
	POSTGRES_DSN string
}

func InitConfigs() *Configs {
	return &Configs{
		POSTGRES_DSN: helpers.MustEnv("POSTGRES_DSN"),
	}
}
