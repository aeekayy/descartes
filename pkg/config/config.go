package config

// AppConfig stores the applications configuration
type AppConfig struct {
	Port string `yaml:"port",json:"port"`
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
