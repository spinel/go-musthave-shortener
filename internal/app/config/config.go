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
	ServerAddress    string `envconfig:"SERVER_ADDRESS"`
	BaseURL          string `envconfig:"BASE_URL"`
	DatabaseDSN      string `envconfig:"DATABASE_DSN"`
	PgMigrationsPath string `envconfig:"PG_MIGRATIONS_PATH"`
	CookieSecretKey  string `envconfig:"COOKIE_SECRET_KEY"`
	BatchQueueSize   int    `envconfig:"BATCH_QUEUE_SIZE"`
}

var (
	config Config
	once   sync.Once
)

const (
	defaultServerAddress    = "localhost:8080"
	defaultBaseURL          = "http://localhost:8080"
	defaultDatabaseDSN      = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	defaultPgMigrationsPath = "file://internal/app/repository/pg/migrations"
	defaultBatchQueueSize   = 10
)

// NewConfig is a singleton env 	config constructor
func NewConfig() Config {
	once.Do(func() {
		err := envconfig.Process("", &config)
		if err != nil {
			log.Fatal(err)
		}

		// bind flags
		if (Config{}) == config {
			bindFlag(config)
		}

		// set default values.
		setDefault(&config)

		// show command line config
		fmt.Println(config)
	})

	return config
}

func setDefault(c *Config) {
	if c.ServerAddress == "" {
		c.ServerAddress = defaultServerAddress
	}
	if c.BaseURL == "" {
		c.BaseURL = defaultBaseURL
	}
	if c.DatabaseDSN == "" {
		c.DatabaseDSN = defaultDatabaseDSN
	}
	if c.PgMigrationsPath == "" {
		c.PgMigrationsPath = defaultPgMigrationsPath
	}
	if c.BatchQueueSize == 0 {
		c.BatchQueueSize = defaultBatchQueueSize
	}
}

func bindFlag(config Config) {
	flag.StringVar(&config.ServerAddress, "a", defaultServerAddress, "app server address")
	flag.StringVar(&config.BaseURL, "b", defaultBaseURL, "base url of links")
	flag.StringVar(&config.DatabaseDSN, "d", defaultDatabaseDSN, "database dsn")
	flag.StringVar(&config.PgMigrationsPath, "m", defaultPgMigrationsPath, "database migrations")
	flag.IntVar(&config.BatchQueueSize, "s", defaultBatchQueueSize, "batch queue size")
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
	if len(c.PgMigrationsPath) == 0 {
		return errors.New("no database migrations")
	}
	if c.BatchQueueSize <= 0 {
		return errors.New("no batch queue size")
	}
	return nil
}
