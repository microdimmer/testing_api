package rest_server

type Config struct {
	DBHost     string `toml:"db_host"`
	DBPort     int    `toml:"db_port"`
	DBName     string `toml:"db_name"`
	DBUser     string `toml:"db_user"`
	DBPassword string `toml:"db_password"`
	BindAddr   string `toml:"bind_addr"`
	LogLevel   string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		DBHost:     "127.0.0.1",
		DBPort:     5432,
		DBName:     "testing_api",
		DBUser:     "postgres",
		DBPassword: "postgres",
		BindAddr:   ":8080",
		LogLevel:   "debug",
	}
}
