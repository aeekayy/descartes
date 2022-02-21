package config

// AppConfig stores the applications configuration
type AppConfig struct {
	Port int `yaml:"port",json:"port"`
}

// Functions
// ******************************
// NewAppConfig return an application configuration
// for use with the API and app
func NewAppConfig(port int) (*AppConfig, error) {
	return &AppConfig{
		Port: port,
	}, nil
}
