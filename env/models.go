package env

type RedisEnv struct {
	Host     string
	Port     string
	Password string
	DB       string
}

type AppEnv struct {
	Port   string
	Domain string
}
