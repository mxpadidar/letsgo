package conf

type Conf struct {
	AppName    string
	AppVersion string
	Debug      bool
}

func NewConf() *Conf {
	loadEnvFile()
	return &Conf{
		AppName:    getEnv("APP_NAME", "MyApp"),
		AppVersion: getEnv("APP_VERSION", "1.0.0"),
		Debug:      getEnv("APP_DEBUG", "false") == "true",
	}
}
