package config

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerAddress string `envconfig:"SERVER_ADDRESS"`
	BaseURL       string `envconfig:"BASE_URL"`
	GobFileName   string `envconfig:"FILE_STORAGE_PATH"`
}

var (
	config Config
	once   sync.Once
)

// NewConfig is a singleton env 	config constructor
func Get() *Config {
	once.Do(func() {
		err := envconfig.Process("", &config)
		if err != nil {
			log.Fatal(err)
		}

		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Configuration:", string(configBytes))
	})

	return &config
}
