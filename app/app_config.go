package app

import "github.com/spf13/viper"

func setupConfig() {
	viper.SetConfigFile("config.env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Warn(err)
		log.Info("Switch to AutomaticEnv")
		viper.AutomaticEnv()
	}
}
