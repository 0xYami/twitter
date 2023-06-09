package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Address     string
	Port        string
	TokenSecret string
	Timeout     time.Duration
	DB          *DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Timeout  time.Duration
}

func (c *DBConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBName,
	)
}

func Load(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	db := &DBConfig{
		Host:     viper.GetString("PGHOST"),
		User:     viper.GetString("PGUSER"),
		Password: viper.GetString("PGPASSWORD"),
		DBName:   viper.GetString("PGDATABASE"),
		Port:     viper.GetString("PGPORT"),
		Timeout:  10 * time.Second,
	}

	return &Config{
		Address:     viper.GetString("ADDRESS"),
		Port:        viper.GetString("PORT"),
		TokenSecret: viper.GetString("TOKEN_SECRET"),
		Timeout:     10 * time.Second,
		DB:          db,
	}, nil
}
