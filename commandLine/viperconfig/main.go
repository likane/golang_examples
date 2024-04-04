package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Init struct {
	Serial string `mapstructure:"serial"`
	Make   string `mapstructure:"make"`
	Model  string `mapstructure:"model"`
}

type TagId struct {
	Id   string `mapstructure:"id"`
	Name string `mapstructure:"name"`
}

type Config struct {
	Init  Init  `mapstructure:"init"`
	TagId TagId `mapstructure:"tagid"`
}

func main() {
	viper.BindEnv("JAVA_HOME")
	val := viper.Get("JAVA_HOME")
	fmt.Println("JAVA_HOME:", val)

	viper.SetConfigType("json")
	viper.SetConfigFile("./session.json")
	fmt.Printf("Using session config: %s\n", viper.ConfigFileUsed())
	viper.ReadInConfig()

	if viper.IsSet("init.serial") {
		fmt.Println("init.serial:", viper.Get("init.serial"))
	} else {
		fmt.Println("init.serial not set.")
	}
	if viper.IsSet("init.make") {
		fmt.Println("init.make:", viper.Get("init.make"))
	} else {
		fmt.Println("init.make is not set.")
	}
	if viper.IsSet("tagId.id") {
		fmt.Println("tagId.id:", viper.Get("tagId.id"))
	} else {
		fmt.Println("tagId.id is not set.")
	}

	viper.SetConfigName("session")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config)
}
