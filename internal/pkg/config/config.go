package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var DEVELOPMENT = "DEVELOPMENT"
var PRODUCTION = "PRODUCTION"

type Config struct {
	Path     string
	FilePath string
	log      *log.Logger
	*viper.Viper
}

func NewConfig() (*Config, error) {
	godotenv.Load(".env")
	viper.AutomaticEnv()
	viper.SetDefault("windows-x", 1433)
	viper.SetDefault("windows-y", 861)
	viper.SetConfigType("json")

	cfgPath, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	cfgPath += "/.popomepost"

	return &Config{
		log:      log.New(os.Stderr, "[CONFIG] ", log.Ldate|log.Ltime|log.Lshortfile),
		Path:     cfgPath,
		FilePath: cfgPath + "/config.json",
		Viper:    viper.GetViper(),
	}, nil

}
func (c *Config) LoadConfig() {
	c.log.Println("running mode " + c.GetString("ENV"))
	if _, err := os.Stat(c.Path); os.IsNotExist(err) {
		c.log.Println("config: create config")
		err = os.Mkdir(c.Path, os.ModeDir|0755)
		if err != nil {
			c.log.Fatalf("config: err %s", err)
		}
		err = viper.WriteConfigAs(c.FilePath)
		if err != nil {
			c.log.Fatalf("config: err %s", err)
		}
	} else {
		c.log.Println("config: load config")
		file, err := os.Open(c.FilePath)
		if err != nil {
			c.log.Fatalf("config: err %s", err)
		}
		viper.ReadConfig(file)
	}
}

func (c *Config) Save() error {
	return viper.WriteConfigAs(c.FilePath)
}
