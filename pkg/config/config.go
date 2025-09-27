package config

import (
	l "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Init load configurations from config.yml file
func Init(cfgFile string, secretFile string) error {
	viper.SetConfigFile(cfgFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		l.Error(err)
		return err
	}

	initConfig()

	return nil
}

// initConfig laod all configurations
func initConfig() {
	loadApp()
	loadDatabase()
	loadGRPCServices()

}
