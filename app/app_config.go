package app

import "github.com/spf13/viper"

var env *Env

type Env struct {
	JwtSecret string
	DbName    string
	DbUser    string
	DbPass    string
	DbPort    string
	DbHost    string
	DbFile    string
}

func setupConfig() {
	viper.SetConfigFile("config.env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Warn(err)
		log.Info("Switch to AutomaticEnv")
		viper.AutomaticEnv()
	}
	env = &Env{
		JwtSecret: mustReadConfig("TODO_JWTSECRET"),
		DbUser:    mustReadConfig("TODO_DBUSER"),
		DbPass:    mustReadConfig("TODO_DBPASS"),
		DbHost:    mustReadConfig("TODO_DBHOST"),
		DbPort:    mustReadConfig("TODO_DBPORT"),
		DbName:    mustReadConfig("TODO_DBNAME"),
		DbFile:    mustReadConfig("TODO_DBFILE"),
	}
}

func mustReadConfig(key string) string {
	str := viper.GetString(key)
	if str == "" {
		log.Fatalf("config %s not found", key)
	}
	log.Infof("config %s found", key)
	return str
}
