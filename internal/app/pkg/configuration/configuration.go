package configuration

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func Load() *Config {
	// TODO: Use .env

	config := &Config{
		Host:     "127.0.0.1",
		Port:     "3306",
		User:     "root",
		Password: "123456",
		Database: "onvet_db",
	}

	return config
}
