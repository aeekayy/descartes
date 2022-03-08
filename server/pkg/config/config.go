package config

import (
	"os"
	"strconv"
)

// AppConfig stores the applications configuration
type AppConfig struct {
	Port string `yaml:"port",json:"port"`
}

// CronConfig stores the cron job configuration
type CronConfig struct {
	DB  CronConfigDB  `yaml:"db",json:"db"`
	API CronConfigAPI `yaml:"api",json:"api"`
}

// CronConfigDB
type CronConfigDB struct {
	Host     string `yaml:"host",json:"host"`
	Port     int    `yaml:"port",json:"port"`
	Name     string `yaml:"name",json:"name"`
	Username string `yaml:"username",json:"username"`
	Password string `yaml:"password",json:"password"`
}

// CronConfigAPI
type CronConfigAPI struct {
	Keys CronConfigAPIKeys `yaml:"keys",json:"keys"`
}

// CronConfigAPIKeys
type CronConfigAPIKeys struct {
	AlphaVantage string `yaml:"alphavantage",json:"alphavantage"`
}

// CronConfigDB
type DBConfig struct {
	Host     string `yaml:"host",json:"host"`
	Port     int    `yaml:"port",json:"port"`
	Name     string `yaml:"name",json:"name"`
	Username string `yaml:"username",json:"username"`
	Password string `yaml:"password",json:"password"`
}

// Functions
// ******************************
// NewAppConfig return an application configuration
// for use with the API and app
func NewAppConfig(port string) (*AppConfig, error) {
	// TODO: Validate that port is a valid integer
	// between 1 and 65535
	return &AppConfig{
		Port: port,
	}, nil
}

// NewCronConfig return a cron job configuration
// from the environment variables
// TODO: Switch to viper
func NewCronConfig() (*CronConfig, error) {
	config := CronConfig{
		API: CronConfigAPI{
			Keys: CronConfigAPIKeys{
				AlphaVantage: "",
			},
		},
		DB: CronConfigDB{
			Name:     "",
			Port:     5432,
			Username: "",
			Password: "",
			Host:     "",
		},
	}
	config.API.Keys.AlphaVantage = os.Getenv("ALPHA_VANTAGE_API_KEY")
	config.DB.Host = os.Getenv("DB_HOST")
	intPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}
	config.DB.Port = intPort
	config.DB.Name = os.Getenv("DB_NAME")
	config.DB.Username = os.Getenv("DB_USERNAME")
	config.DB.Password = os.Getenv("DB_PASSWORD")

	return &config, nil
}
