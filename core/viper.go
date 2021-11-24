package core

import "github.com/spf13/viper"


func initViper() {
	viper.AddConfigPath("./config")
	viper.ReadInConfig()
}
