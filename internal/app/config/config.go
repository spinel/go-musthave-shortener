package config

import (
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

const (
	defaultServerAddress = "localhost:8080"
	defaultBaseURL       = "http://localhost:8080"
	defaultCobFileName   = "urls.gob"
)

// NewConfig is a singleton env 	config constructor
func NewConfig() *Config {
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

		// show command line config
		fmt.Println(config)
	})

	return &config
}

func bindFlag(config *Config) {
	flag.StringVar(&config.ServerAddress, "a", defaultServerAddress, "app server address")
	flag.StringVar(&config.BaseURL, "b", defaultBaseURL, "base url of links")
	flag.StringVar(&config.GobFileName, "f", defaultCobFileName, "gob file path")
	flag.Parse()
}
