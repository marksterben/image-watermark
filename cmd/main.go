package main

import (
	"image-watermark/server"
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`.env`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`DEBUG`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	server := server.NewServer()
	server.Run()
}
