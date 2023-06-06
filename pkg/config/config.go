package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func NewAppConfig(fileName string) error {
	viper.SetConfigName(fileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func GetConfigString(key Key) string {
	return viper.GetString(fmt.Sprint(key))
}

func GetConfigInt(key Key) int {
	return viper.GetInt(fmt.Sprint(key))
}

func GetConfigBool(key Key) bool {
	return viper.GetBool(fmt.Sprint(key))
}
