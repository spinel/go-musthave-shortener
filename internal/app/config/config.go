package config

import (
	"encoding/json"
	"flag"
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

const defaultServerAddress = "localhost:9080"
const defaultBaseURL = "http://localhost:9080"
const defaultCobFileName = "urls.gob"

// NewConfig is a singleton env 	config constructor
func Get() *Config {
	once.Do(func() {
		err := envconfig.Process("", &config)
		if err != nil {
			log.Fatal(err)
		}

		// bind flags or default
		// values if env is empty
		if (Config{}) == config {
			bindFlag(&config)
		}

		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Configuration:", string(configBytes))
	})

	return &config
}

func bindFlag(config *Config) {
	flag.StringVar(&config.ServerAddress, "a", defaultServerAddress, "app server address")
	flag.StringVar(&config.BaseURL, "b", defaultBaseURL, "base url of links")
	flag.StringVar(&config.GobFileName, "f", defaultCobFileName, "gob file path")
	flag.Parse()
}
