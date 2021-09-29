package goconfigcloud

import (
	"strings"

	"github.com/spf13/viper"
)

func EnvironmentVariable() {
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

}
