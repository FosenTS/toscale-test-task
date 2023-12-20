package config

import (
	"github.com/spf13/viper"
	"time"
)

const gormConfigFileName = "gorm.config.yaml"

type GormConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
	SSLMode      string
	TimeZone     *time.Location

	MaxIdleConnections    int
	MaxOpenConnection     int
	ConnectionMaxLifeTime time.Duration
}

var gormConfigInst = &GormConfig{}

func Gorm() (*GormConfig, error) {
	gormConfigInst.Host = viper.GetString("host")
	gormConfigInst.Port = viper.GetInt("port")
	gormConfigInst.User = viper.GetString("username")
	gormConfigInst.Password = viper.GetString("password")
	gormConfigInst.DatabaseName = viper.GetString("dbName")
	gormConfigInst.SSLMode = viper.GetString("sslMode")
	timeZone, err := time.LoadLocation(viper.GetString("timeZone"))
	if err != nil {
		return nil, err
	}

	gormConfigInst.TimeZone = timeZone

	gormConfigViper := viper.New()
	gormConfigViper.AddConfigPath(viper.GetString("configPath"))
	gormConfigViper.SetConfigName(gormConfigFileName)
	gormConfigViper.SetConfigType("yaml")
	err = gormConfigViper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	gormConfigInst.MaxIdleConnections = gormConfigViper.GetInt("maxIdleConnections")
	gormConfigInst.MaxOpenConnection = gormConfigViper.GetInt("maxOpenConnections")
	gormConfigInst.ConnectionMaxLifeTime = time.Duration(gormConfigViper.GetInt("connectionMaxLifeTimeSecond")) * time.Second

	return gormConfigInst, nil
}
