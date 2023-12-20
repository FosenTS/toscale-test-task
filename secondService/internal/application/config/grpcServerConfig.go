package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type GRPCConfig struct {
	host string
	port int

	Address string
}

var gRPCConfig = &GRPCConfig{}

func GRPC() (*GRPCConfig, error) {
	gRPCConfig.host = viper.GetString("gRPCHost")
	gRPCConfig.port = viper.GetInt("gRPCPort")

	gRPCConfig.Address = fmt.Sprintf("%s:%d", gRPCConfig.host, gRPCConfig.port)

	return gRPCConfig, nil
}
