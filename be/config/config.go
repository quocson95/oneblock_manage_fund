package config

import (
	"io"

	"gopkg.in/yaml.v3"
)

type ConfigSet struct {
	Be     Be            `yaml:"be"`
	DB     DBConfig      `yaml:"db"`
	Google GoogleConsole `yaml:"google"`
}
type Be struct {
	Port int `yaml:"port"`
}
type DBConfig struct {
	DSN string `yaml:"dsn"`
	// Schema string `yaml:"schema"`
}

type GoogleConsole struct {
	ID          string `yaml:"id"`
	Secret      string `yaml:"secret"`
	CallbackSSO string `yaml:"callback_sso"`
	RedirectURI string `yaml:"redirect_uri"`
}

var defaultConfig = &ConfigSet{}

func Parse(reader io.Reader) error {
	err := yaml.NewDecoder(reader).Decode(defaultConfig)
	if err != nil {
		return err
	}
	return nil
}

func Get() *ConfigSet {
	return defaultConfig
}
