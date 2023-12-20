package config

import "github.com/spf13/viper"

const binanceFilename = "binance.config.yaml"

type BinanceConfig struct {
	Domain    string
	KlinePath string
}

var binanceConfig = &BinanceConfig{}

func Binance() (*BinanceConfig, error) {
	binanceConfigViper := viper.New()
	binanceConfigViper.AddConfigPath(viper.GetString("configPath"))
	binanceConfigViper.SetConfigName(binanceFilename)
	binanceConfigViper.SetConfigType("yaml")
	err := binanceConfigViper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	binanceConfig.Domain = binanceConfigViper.GetString("domain")
	binanceConfig.KlinePath = binanceConfigViper.GetString("klinePath")

	return binanceConfig, nil
}
