package config

import (
	"github.com/spf13/viper"
)

var v *viper.Viper

func Get() *viper.Viper {
	return v
}

func InjectConfig(cfg *viper.Viper) {
	v = cfg
}
