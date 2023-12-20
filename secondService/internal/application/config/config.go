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

	err = viper.BindEnv("configPath", "CONFIG_PATH")
	if err != nil {
		return err
	}

	err = viper.BindEnv("gRPCHost", "GRPC_HOST")
	if err != nil {
		return err
	}

	err = viper.BindEnv("gRPCPort", "GRPC_PORT")
	if err != nil {
		return err
	}

	return nil
}
