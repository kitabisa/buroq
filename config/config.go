package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
}

var defaultConfig *viper.Viper

func Config() Provider {
	return defaultConfig
}

func init() {
	defaultConfig = readViperConfig()
}

func readViperConfig() *viper.Viper {
	v := viper.New()
	v.AddConfigPath(".")
	v.AddConfigPath("./params")
	v.AddConfigPath("/opt/go-bootstrap/params")
	v.SetConfigName("go-bootstrap")

	err := v.ReadInConfig()
	if err == nil {
		fmt.Printf("Using config file: %s \n\n", v.ConfigFileUsed())
	} else {
		panic(fmt.Errorf("Config error: %s", err))
	}

	return v
}
