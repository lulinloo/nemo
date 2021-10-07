package config

import (
	"errors"
	"log"
	"os"
	"path"

	"github.com/Shopify/sarama"
	"github.com/adrg/xdg"
	"github.com/spf13/viper"
)

var defaultConfig = &Config{
	Dbtypes: map[string]*Instance{
		"mysql": {
			Gap: 2,
		},
	},
}

type Config struct {
	filepath    string
	DefaultType string
	Dbtypes     map[string]*Instance
}

func (c *Config) Filepath() string {
	return c.filepath
}

type Instance struct {
	Gap int
}

func (c *Config) Instance(profile string) (int, *sarama.Config, error) {
	if profile == "" {
		profile = c.DefaultType
	}
	prof, ok := c.Dbtypes[profile]
	if !ok {
		log.Fatal("No profile specified")
	}

	cfg := sarama.NewConfig()

	return prof.Gap, cfg, nil
}

func LoadConfig(cfgFilepath string) (*Config, error) {
	cfgRoot := path.Join(xdg.ConfigHome, "nemo")
	if _, err := os.Stat(cfgRoot); os.IsNotExist(err) {
		if err := os.Mkdir(cfgRoot, 0755); err != nil {
			log.Fatal(err)
		}
	}

	if cfgFilepath != "" {
		viper.SetConfigFile(cfgFilepath)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath(cfgRoot)
	}

	if err := viper.ReadInConfig(); err != nil {
		var errNotFound viper.ConfigFileNotFoundError
		if cfgFilepath == "" && errors.As(err, &errNotFound) {
			return defaultConfig, nil
		} else {
			log.Fatal(err)
		}
	}

	cfg := &Config{
		filepath: viper.GetViper().ConfigFileUsed(),
	}
	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatal(err)
	}

	return cfg, nil
}
