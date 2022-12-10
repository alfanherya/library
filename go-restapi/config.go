package gorestapi

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var ErrorLoadConfig error

func InitConfig() {
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(true)

}

func HttpEndpoint() string {
	v := viper.GetString("httpEndpoint")
	if v == "" {
		return ":8080"
	}
	return v
}

func SkipAuditLoggerURIPattern() string {
	v := viper.GetString("logger.audit.uri.skip.regex")
	if v == "" {
		return "^/swagger/*|^/metrics.*s"
	}
	return v
}
