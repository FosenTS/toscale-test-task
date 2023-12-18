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

func Gorm() (GormConfig, error) {
	gormConfigInst.Host = viper.GetString("host")
	gormConfigInst.Port = viper.GetInt("port")
	gormConfigInst.User = viper.GetString("username")
	gormConfigInst.Password = viper.GetString("password")
	gormConfigInst.DatabaseName = viper.GetString("dbName")
	gormConfigInst.SSLMode = viper.GetString("sslMode")
	timeZone, err := time.LoadLocation(viper.GetString("timeZone"))
	if err != nil {
		return GormConfig{}, err
	}

	gormConfigInst.TimeZone = timeZone

	viper.SetConfigFile(gormConfigFileName)
	viper.SetConfigType("yaml")
	gormConfigInst.MaxIdleConnections = viper.GetInt("maxIdleConnections")
	gormConfigInst.MaxOpenConnection = viper.GetInt("maxOpenConnections")
	gormConfigInst.ConnectionMaxLifeTime = time.Duration(viper.GetInt("connectionMaxLifeTimeSecond")) * time.Second

	return *gormConfigInst, nil
}
