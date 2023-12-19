package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const envFileName = ".env.local"

func Env() error {
	err := godotenv.Load(envFileName)
	if err != nil {
		return err
	}

	err = viper.BindEnv("projectPath", "PROJECT_PATH")
	if err != nil {
		return err
	}
	err = viper.BindEnv("configPath", "CONFIG_PATH")
	if err != nil {
		return err
	}

	err = viper.BindEnv("httpHost", "HTTP_HOST")
	if err != nil {
		return err
	}

	err = viper.BindEnv("httpPort", "HTTP_PORT")
	if err != nil {
		return err
	}

	err = viper.BindEnv("host", "HOST")
	if err != nil {
		return err
	}
	err = viper.BindEnv("port", "PORT")
	if err != nil {
		return err
	}
	err = viper.BindEnv("username", "USERNAME")
	if err != nil {
		return err
	}
	err = viper.BindEnv("password", "PASSWORD")
	if err != nil {
		return err
	}
	err = viper.BindEnv("dbName", "DB_NAME")
	if err != nil {
		return err
	}
	err = viper.BindEnv("sslMode", "SSL_MODE")
	if err != nil {
		return err
	}
	err = viper.BindEnv("timeZone", "TIME_ZONE")
	if err != nil {
		return err
	}

	return nil
}
