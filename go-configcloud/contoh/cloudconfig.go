package main

import (
	"fmt"

	go_config "github.com/alfanherya/library/go-configcloud"
	"github.com/spf13/viper"
)

func main() {
	go_config.SpringCloudConfig("cloud", "https://spring-cloud-server.herokuapp.com/goswagger/development/master")
	fmt.Println(viper.Get("cloud.appname"))
	fmt.Println(viper.Get("cloud.port"))
	fmt.Println(viper.Get("cloud.databasename"))
	fmt.Println(viper.AllKeys())

	go_config.EnvironmentVariable()
}
