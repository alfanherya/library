package goconfigcloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type springCloudConfig struct {
	Name            string           `json:"name"`
	Profiles        []string         `json:"profiles"`
	Label           interface{}      `json:"label"`
	Version         string           `json:"version"`
	State           interface{}      `json:"state"`
	PropertySources []propertySource `json:"propertySources"`
}
type propertySource struct {
	Name   string                 `json:"name"`
	Source map[string]interface{} `json:"source"`
}

func SpringCloudConfig(prefix, url string) {
	body, err := callSpringCloudConfig(url)
	if err != nil {
		logrus.Printf("SpringCloudConfig: %s\n", err)
		return
	}
	clconf := new(springCloudConfig)
	err = json.Unmarshal(body, clconf)
	if err != nil {
		logrus.Printf("SpringCloudConfig: %s\n", err)
		return
	}
	for _, vps := range clconf.PropertySources {
		for is, vs := range vps.Source {
			if prefix == "" {
				viper.Set(fmt.Sprintf("%s", is), vs)
			} else {
				viper.Set(fmt.Sprintf("%s.%s", prefix, is), vs)
			}

		}
	}
}

func callSpringCloudConfig(url string) ([]byte, error) {
	method := "GET"
	client := &http.Client{
		Timeout: 120 * time.Second,
	}
	trx, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("go-config error preparing http request to %s with error: %s", url, err.Error())
	}
	trx.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(trx)
	if err != nil {
		return nil, fmt.Errorf("go-config error services %s with error: %s", url, err)
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
