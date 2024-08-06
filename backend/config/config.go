package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server     ServerConfig
	GoogleAuth GoogleAuthConfig
}

type ServerConfig struct {
	Port string
}

type GoogleAuthConfig struct {
	ClientID     string
	ClientSecret string
	CallbackURL  string
}

func GetConfig() *Config {
	v, err := LoadConfig("config", "yml")
	if err != nil {
		log.Fatalf("Error in load config %v", err)
	}

	cfg, err := ParseConfig(v)
	if err != nil {
		log.Fatalf("Error in parse config %v", err)
	}
	cfg.Server = *loadServerConfig(v)
	cfg.GoogleAuth = *loadGoogleAuthConfig(v)

	return cfg
}

func loadServerConfig(v *viper.Viper) *ServerConfig {
	var servConf ServerConfig
	servConf.Port = (v.Get("SERVER.PORT")).(string)
	return &servConf
}

func loadGoogleAuthConfig(v *viper.Viper) *GoogleAuthConfig {
	var googleAuthConfig GoogleAuthConfig
	googleAuthConfig.ClientID = (v.Get("GOOGLE.CLIENT_ID")).(string)
	googleAuthConfig.ClientSecret = (v.Get("GOOGLE.CLIENT_SECRET")).(string)
	googleAuthConfig.CallbackURL = (v.Get("GOOGLE.CALLBACK_URL")).(string)
	return &googleAuthConfig
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to parse config: %v", err)
		return nil, err
	}
	return &cfg, nil
}

func LoadConfig(filename string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(filename)
	v.AddConfigPath("./config")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("Unable to read config: %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}
