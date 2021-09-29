package goconfigcloud

import (
	cf "github.com/cloudfoundry-community/go-cfenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func PCFCups() {
	vcap, _ := cf.Current()
	vcapServices, _ := vcap.Services.WithNameUsingPattern(".*")
	if cf.IsRunningOnCF() {
		for _, vcapServices := range vcapServices {
			logrus.Println("load vcap services: ", vcapServices.Name)
			viper.Set(vcapServices.Name, vcapServices.Credentials)
		}
	}
}
