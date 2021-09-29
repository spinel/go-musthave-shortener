package config

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/url"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerAddress string `envconfig:"SERVER_ADDRESS"`
	BaseURL       string `envconfig:"BASE_URL"`
	DatabaseDSN   string `envconfig:"DATABASE_DSN"`
}

var (
	config Config
	once   sync.Once
)

const (
	defaultServerAddress = "localhost:8080"
	defaultBaseURL       = "http://localhost:8080"
	defaultDatabaseDSN   = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
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
	flag.StringVar(&config.DatabaseDSN, "d", defaultDatabaseDSN, "database dsn")
	flag.Parse()
}

func (c Config) Validate() error {
	if (Config{}) == c {
		return errors.New("empty config")
	}
	_, err := url.ParseRequestURI(c.BaseURL)
	if err != nil {
		return err
	}
	_, err = url.ParseRequestURI(c.ServerAddress)
	if err != nil {
		return err
	}
	if len(c.DatabaseDSN) == 0 {
		return errors.New("no database dsn")
	}
	return nil
}
