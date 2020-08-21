package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

const appName = "ss" // supersense

type Config struct {
	Port int64 `default:"4000" split_words:"true"`

	GithubToken string   `split_words:"true"`
	GithubRepos []string `split_words:"true"`

	DummyPeriod  string `split_words:"true"`
	DummyMessage string `split_words:"true"`

	TwitterConsumerKey    string   `split_words:"true"`
	TwitterConsumerSecret string   `split_words:"true"`
	TwitterAccessToken    string   `split_words:"true"`
	TwitterAccessSecret   string   `split_words:"true"`
	TwitterQuery          []string `split_words:"true"`

	GraphQLPlayground bool `envconfig:"GRAPHQL_PLAYGROUND" default:"false"`

	Persistence bool `split_words:"true" default:"false"`
	PersistenceBoltDBFilePath string `envconfig:"PERSISTENCE_BOLTDB_FILEPATH" default:"ss.db"`
	PersistenceRedisAddress string `split_words:"true"`
	PersistenceRedisPassword string `split_words:"true"`
}

func load(appName string) (*Config, error) {
	conf := new(Config)
	if err := envconfig.Process(appName, conf); err != nil {
		return nil, errors.WithStack(err)
	}
	return conf, nil
}

// LoadDefault load a minimal default configuration
func LoadDefault() (*Config, error) {
	return load(appName)
}