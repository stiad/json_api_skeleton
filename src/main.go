package main

import (
	"github.com/spf13/viper"
	"github.com/stiad/json_api_skeleton/src/app"
	"log"
)

func init() {
	viper.SetConfigFile("example.env")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Error reading config file, ", err)
	}
	viper.WatchConfig()
}

func main() {
	app.NewServer().Serve(":8081")
}
