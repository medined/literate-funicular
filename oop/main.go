package main

import (
	"fmt"
	"github.com/spf13/viper"
	"oop/configuration"
)

func main() {
	viper.SetConfigName("literate-funicular.application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	c := configuration.New(viper.GetString("GITHUB_ACCESS_TOKEN"), viper.GetString("WORKDIR"))
	fmt.Println(c)
}
