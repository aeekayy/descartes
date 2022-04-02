package mqtt

import (
	"errors"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Config struct {
	Host     string            `yaml:"host",json:"host"`
	Port     int               `yaml:"port",json:"port"`
	Username string            `yaml:"username",json:"username"`
	Password string            `yaml:"password",json:"password"`
	Topics   map[string]string `yaml:"topics",json:"topics"`
}

func NewClient(clientID string, conf *Config) (mqtt.Client, error) {
	// this should really be a web hook but for now, we'll just use mqtt
	// for quick and dirty
	if conf.Username == "" || conf.Password == "" {
		return nil, errors.New("invalid mqtt credential")
	}
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", conf.Host, conf.Port))
	opts.SetClientID(clientID)      // set a name as you desire
	opts.SetUsername(conf.Username) // these are the credentials that you declare for your cluster
	opts.SetPassword(conf.Password)

	client := mqtt.NewClient(opts)

	return client, nil
}
