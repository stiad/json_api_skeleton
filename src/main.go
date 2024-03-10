package main

import (
	"github.com/spf13/viper"
	"github.com/stiad/json_api_skeleton/src/app"
	"log"
)

func init() {
	viper.SetConfigFile("example.env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error reading config file, ", err)
	}
	viper.WatchConfig()
}

func main() {
	server := app.NewServer()
	server.Serve(":8081")
}
